# tree-sitter-c_with_esql

C grammar with Embedded SQL support for [tree-sitter](https://github.com/tree-sitter/tree-sitter).  
Forked from [tree-sitter-c](https://github.com/tree-sitter/tree-sitter-c).  
Adapted from [this C99 grammar](http://slps.github.io/zoo/c/iso-9899-tc3.html).  

Documentation for Embedded SQL:

- [Embedded SQL - Wikipedia](https://en.wikipedia.org/wiki/Embedded_SQL)
- [Pro*C - Wikipedia](https://en.wikipedia.org/wiki/Pro*C)
- [Introduction to Pro*C Embedded SQL](http://infolab.stanford.edu/%7Eullman/fcdb/oracle/or-proc.html)
- [Oracle  Pro*C Sample Programs](https://download.oracle.com/otn_hosted_doc/timesten/1122/quickstart/html/developer/proc/proc.html)

## NeoVim configuration for lazy.nvim

```lua
-- nvim/lua/plugins/tree-sitter-c_with_esql.lua:
---@module "lazy"

return {
  "sealor/tree-sitter-c_with_esql",
  dependencies = { "nvim-treesitter/nvim-treesitter" },

  ---@param plugin LazyPlugin
  config = function(plugin, _)
    local parser_configs = require("nvim-treesitter.parsers").get_parser_configs()

    ---@type ParserInfo
    local c_with_esql = {
      install_info = {
        url = plugin.dir,
        files = { "src/parser.c" },
        generate_requires_npm = false,
        requires_generate_from_grammar = false,
      },
      maintainers = { "@sealor" },
      filetype = "c_with_esql",
    }

    parser_configs["c_with_esql"] = c_with_esql

    vim.filetype.add({ extension = { pc = "c_with_esql" } })
  end,

  ---@param plugin LazyPlugin
  build = function(plugin)
    plugin.config(plugin, {})
    require("nvim-treesitter.install").update()({ "c_with_esql" })
  end,
}
```
