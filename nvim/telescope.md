# Configuration

```
local M = {
  "nvim-telescope/telescope.nvim",
  commit = "203bf5609137600d73e8ed82703d6b0e320a5f36",
  event = "Bufenter",
  cmd = { "Telescope" },
  dependencies = {
    {
      "ahmedkhalf/project.nvim",
      commit = "685bc8e3890d2feb07ccf919522c97f7d33b94e4",
    },
  },
}

local actions = require "telescope.actions"

M.opts = {
  defaults = {
    prompt_prefix = " ",
    selection_caret = " ",
    path_display = { "smart" },
    file_ignore_patterns = { ".git/", "node_modules" },
    mappings = {
      i = {
        ["<Down>"] = actions.move_selection_next,
        ["<Up>"] = actions.move_selection_previous,
        ["<C-j>"] = actions.move_selection_next,
        ["<C-k>"] = actions.move_selection_previous,
      },
    },

      planets = {
        show_pluto = true,
        show_moon = true,
      },
      git_files = {
        hidden = true,
        show_untracked = true,
      },
      colorscheme = {
        enable_preview = true,
      },
    },
    pickers = {
      find_files = {
        theme = "dropdown",
        hidden = true,
      },
      live_grep = {
        --@usage don't include the filename in the search results
        only_sort_text = true,
      },
      grep_string = {
        only_sort_text = true,
      },
      buffers = {
        theme = "dropdown",
        previewer = false,
        -- initial_mode = "normal",
        mappings = {
          i = {
            ["<C-d>"] = actions.delete_buffer,
          },
          n = {
            ["dd"] = actions.delete_buffer,
          },
        },
      },
    extensions = {
      fzf = {
        fuzzy = true, -- false will only do exact matching
        override_generic_sorter = true, -- override the generic sorter
        override_file_sorter = true, -- override the file sorter
        case_mode = "smart_case", -- or "ignore_case" or "respect_case"
      },
    },
  },
}

return M

```
