local appy = require("appy")

local plugin = {}

function plugin:on_configure()
    -- Create standard directories for the provider
    appy.mkdir(self.provider_root .. "macros")
    appy.mkdir(self.provider_root .. "queries")
    appy.mkdir(self.provider_root .. "migrations")
    appy.mkdir(self.provider_root .. "driver")

    -- Copy over template files
    appy.copy_file(self.script_root .. "templates/query.sql", self.provider_root .. "queries/query.sql")
    appy.copy_file(self.script_root .. "templates/macro.sql", self.provider_root .. "macros/macro.sql")
    appy.copy_file(self.script_root .. "templates/migration.sql", self.provider_root .. "migrations/migration.sql")
    appy.copy_file(self.script_root .. "templates/core.sql", self.provider_root .. "migrations/core.sql")
end

function plugin:on_load()
    local regenerate_code = function(file, op)
        local project_name = appy.get_project_name()
        project_name = project_name:gsub("%s+", "_")

        print("Regenerating code for file: " .. file .. " with operation: " .. op)
        appy.execute_shell("/usr/bin/python3", {
            self.script_root .. "tools/query_engine.py",
            self.provider_root,
            project_name
        })
    end

    -- Attach watchers
    appy.watch_directory(self.provider_root .. "macros", regenerate_code)
    appy.watch_directory(self.provider_root .. "queries", regenerate_code)
    appy.watch_directory(self.provider_root .. "migrations", regenerate_code)
end

return plugin
