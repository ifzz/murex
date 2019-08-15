package docs

func init() {

	Definition["config"] = "# _murex_ Language Guide\n\n## Command Reference: `config`\n\n> Query or define _murex_ runtime settings\n\n### Description\n\nRather than _murex_ runtime settings being definable via obscure environmental\nvariables, _murex_ instead supports a registry of config defined via the\n`config` command. This means any preferences and/or runtime config becomes\ncentralised and discoverable.\n\n### Usage\n\nList all settings\n\n    config -> <stdout>\n    \nGet a setting\n\n    config get app key -> <stdout>\n    \nSet a setting\n\n    config set app key value\n    \n    <stdin> -> config set app key\n    \n    config eval app key { code-block }\n    \nDefine a new config setting\n\n    config define app key { mxjson }\n    \nReset a setting to it's default value\n\n    !config app key\n    \n    config default app key\n\n### Detail\n\nWith regards to `config`, the following terms are applied:\n\n#### \"app\"\n\nThis refers to a grouped category of settings. For example the name of a built\nin.\n  \nOther app names include\n\n* `shell`: for \"global\" (system wide) _murex_ settings\n* `proc`: for scoped _murex_ settings\n* `http`: for settings that are applied to any processes which use the builtin\n   HTTP user agent (eg `open`, `get`, `getfile`, `set`)\n\n#### \"key\"\n\nThis refers to the config setting itself. For example the \"app\" might be `http`\nbut the \"key\" might be `timeout` - where the \"key\", in this instance, holds the\nvalue for how long any HTTP user agents might wait before timing out.\n\n#### \"value\"\n\nValue is the actual value of a setting. So the value for \"app\": `http`, \"key\":\n`timeout` might be `10`. eg\n\n    » config get http timeout\n    10\n    \n#### \"scope\" / \"scoped\"\n\nSettings in `config`, by default, are scoped per function and module. Any\nfunctions called will inherit the settings of it's caller parent. However any\nchild functions that then change the settings will only change settings for it's\nown function and not the parent caller.\n\nPlease note that `config` settings are scoped differently to local variables.\n\n#### \"global\"\n\nGlobal settings defined inside a function will affect settings queried inside\nanother executing function (same concept as global variables).\n### Directives\n\nThe directives for `config define` are listed below. Headings are formatted\nas follows: \n\n    \"DirectiveName\": json data-type (default value)\n    \nWhere \"default value\" is what will be auto-populated if you don't include that\ndirective (or \"required\" if the directive must be included).\n\n#### \"DataType\": string (required)\n\nThis is the _murex_ data-type for the value.\n\n#### \"Description\": string (required)\n\nDescription is a required field to force developers into writing meaning hints\nenabling the discoverability of settings within _murex_.\n\n#### \"Global\": boolean (false)\n\nThis defines whether this setting is global or scoped.\n\nAll **Dynamic** settings _must_ also be **Global**. This is because **Dynamic**\nsettings rely on a state that likely isn't scoped (eg the contents of a config\nfile).\n\n#### \"Default\": any (required)\n\nThis is the initialized and default value.\n\n#### \"Options\": array (nil)\n\nSome suggested options (if known) to provide as autocompletion suggestions in\nthe interactive command line.\n\n#### \"Dynamic\": map of strings (nil)\n\nOnly use this if config options need to be more than just static values stored\ninside _murex_'s runtime. Using **Dynamic** means `autocomplete get app key`\nand `autocomplete set app key value` will spawn off a subshell running a code\nblock defined from the `Read` and `Write` mapped values. eg\n\n    # Create the example config file\n    (this is the default value) -> > example.conf\n    \n    # mxjson format, so we can have comments and block quotes: #, (, )\n    config define example test ({\n        \"Description\": \"This is only an example\",\n        \"DataType\": \"str\",\n        \"Global\": true,\n        \"Dynamic\": {\n            \"Read\": ({\n                open example.conf\n            }),\n            \"Write\": ({\n                -> > example.conf\n            })\n        },\n        # read the config file to get the default value\n        \"Default\": \"${open example.conf}\"\n    })\n    \nIt's also worth noting the different syntax between **Read** and **Default**.\nThe **Read** code block is being executed when the **Read** directive is being\nrequested, whereas the **Default** code block is being executed when the JSON\nis being read.\n\nIn technical terms, the **Default** code block is being executed by _murex_ \nwhen `config define` is getting executed where as the **Read** and **Write**\ncode blocks are getting stored as a JSON string and then executed only when\nthose hooks are getting triggered.\n\nSee the `mxjson` data-type for more details.\n\n#### \"Dynamic\": { \"Read\": string (\"\") }\n\nThis is executed when `autocomplete get app key` is ran. The STDOUT of the code\nblock is the setting's value.\n\n#### \"Dynamic\": { \"Write\": string (\"\") }\n\nThis is executed when `autocomplete` is setting a value (eg `set`, `default`,\n`eval`). is ran. The STDIN of the code block is the new value.\n\n### See Also\n\n* [`runtime`](../commands/runtime.md):\n  Returns runtime information on the internal state of _murex_\n* [events](../commands/events.md):\n  \n* [mxjson](../types/mxjson.md):\n  Murex-flavoured JSON (primitive)\n* [open](../commands/open.md):\n  "

}
