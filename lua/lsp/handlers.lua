local icons = require("icons")
local M = {}

M.capabilities = vim.lsp.protocol.make_client_capabilities()

local cmp_nvim_lsp = require("cmp_nvim_lsp")
M.capabilities.textDocument.completion.completionItem.snippetSupport = true
M.capabilities = cmp_nvim_lsp.default_capabilities(M.capabilities)

M.setup = function()
	local signs = {
		{ name = "DiagnosticSignError", text = icons.diagnostics.Error },
		{ name = "DiagnosticSignWarn", text = icons.diagnostics.Warning },
		{ name = "DiagnosticSignHint", text = icons.diagnostics.Hint },
		{ name = "DiagnosticSignInfo", text = icons.diagnostics.Information },
	}

	for _, sign in ipairs(signs) do
		vim.fn.sign_define(sign.name, { texthl = sign.name, text = sign.text, numhl = "" })
	end

	local config = {
		-- disable virtual text
		virtual_lines = true,
		virtual_text = false,
		-- virtual_text = {
		--   -- spacing = 7,
		--   -- update_in_insert = false,
		--   -- severity_sort = true,
		--   -- prefix = "<-",
		--   prefix = " ●",
		--   source = "if_many", -- Or "always"
		--   -- format = function(diag)
		--   --   return diag.message .. "blah"
		--   -- end,
		-- },

		-- show signs
		signs = {
			active = signs,
		},
		update_in_insert = true,
		underline = true,
		severity_sort = true,
		float = {
			focusable = true,
			style = "minimal",
			border = "rounded",
			-- border = {"▄","▄","▄","█","▀","▀","▀","█"},
			source = "if_many", -- Or "always"
			header = "",
			prefix = "",
			-- width = 40,
		},
	}

	vim.diagnostic.config(config)

	-- Function to check if a floating dialog exists and if not
	-- then check for diagnostics under the cursor
	function OpenDiagnosticIfNoFloat()
		for _, winid in pairs(vim.api.nvim_tabpage_list_wins(0)) do
			if vim.api.nvim_win_get_config(winid).zindex then
				return
			end
		end
		-- THIS IS FOR BUILTIN LSP
		vim.diagnostic.open_float(0, {
			scope = "cursor",
			focusable = false,
			close_events = {
				"CursorMoved",
				"CursorMovedI",
				"BufHidden",
				"InsertCharPre",
				"WinLeave",
			},
		})
	end
	-- Show diagnostics under the cursor when holding position
	vim.api.nvim_create_augroup("lsp_diagnostics_hold", { clear = true })
	vim.api.nvim_create_autocmd({ "CursorHold" }, {
		pattern = "*",
		command = "lua OpenDiagnosticIfNoFloat()",
		group = "lsp_diagnostics_hold",
	})

	vim.lsp.handlers["textDocument/hover"] = vim.lsp.with(vim.lsp.handlers.hover, {
		border = "rounded",
		-- width = 60,
		-- height = 30,
	})

	vim.lsp.handlers["textDocument/signatureHelp"] = vim.lsp.with(vim.lsp.handlers.signature_help, {
		border = "rounded",
		-- width = 60,
		-- height = 30,
	})
end

local function lsp_keymaps(bufnr)
	local opts = { noremap = true, silent = true }
	vim.api.nvim_buf_set_keymap(bufnr, "n", "gd", "<cmd>Telescope lsp_definitions<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "gD", "<cmd>Telescope lsp_declarations<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "gI", "<cmd>Telescope lsp_implementations<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "gr", "<cmd>Telescope lsp_references<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "gl", "<cmd>lua vim.diagnostic.open_float()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "ca", "<cmd>lua vim.lsp.buf.code_action()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "<leader>rn", "<cmd>lua vim.lsp.buf.rename()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "H", "<cmd>lua vim.lsp.buf.hover()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "i", "<C-s>", "<cmd>lua vim.lsp.buf.signature_help()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "<leader>sh", "<cmd>lua vim.lsp.buf.signature_help()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "[d", '<cmd>lua vim.diagnostic.goto_prev({ border = "rounded" })<CR>', opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "]d", '<cmd>lua vim.diagnostic.goto_next({ border = "rounded" })<CR>', opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "<leader>q", "<cmd>lua vim.diagnostic.setloclist()<CR>", opts)
	vim.api.nvim_buf_set_keymap(bufnr, "n", "<leader>cl", "<cmd>lua vim.lsp.codelens.refresh()<CR>", opts)
end

M.on_attach = function(client, bufnr)
	lsp_keymaps(bufnr)

	if client.name == "jdtls" then
		require("jdtls").setup_dap({ hotcodereplace = "auto" })
		require("jdtls.dap").setup_dap_main_class_configs()
	end
	client.server_capabilities.document_formatting = true
end

function M.enable_format_on_save()
	vim.cmd([[
    augroup format_on_save
      autocmd! 
      autocmd BufWritePre * lua vim.lsp.buf.format({ async = false }) 
    augroup end
  ]])
	vim.notify("Enabled format on save")
end

function M.disable_format_on_save()
	M.remove_augroup("format_on_save")
	vim.notify("Disabled format on save")
end

function M.toggle_format_on_save()
	if vim.fn.exists("#format_on_save#BufWritePre") == 0 then
		M.enable_format_on_save()
	else
		M.disable_format_on_save()
	end
end

function M.remove_augroup(name)
	if vim.fn.exists("#" .. name) == 1 then
		vim.cmd("au! " .. name)
	end
end

vim.cmd([[ command! LspToggleAutoFormat execute 'lua require("lsp.handlers").toggle_format_on_save()' ]])

return M