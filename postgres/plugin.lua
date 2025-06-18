local appy = require("appy")

local plugin = {}

function plugin:__generate_driver()
    local project_name = appy.config.project
    project_name = project_name:gsub("-", "_")

    appy.execute_shell("/usr/bin/python3", {
        self.script_root .. "tools/query_engine.py",
        self.provider_root,
        project_name
    })
end

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

    -- Generate the driver code
    plugin:__generate_driver()
end

function plugin:on_load()
    local regenerate_code = function(file, op)
        print("Regenerating code for file: " .. file .. " with operation: " .. op)
        plugin:__generate_driver()
    end

    -- Attach watchers
    appy.watch_directory(self.provider_root .. "macros", regenerate_code)
    appy.watch_directory(self.provider_root .. "queries", regenerate_code)
    appy.watch_directory(self.provider_root .. "migrations", regenerate_code)
end

function plugin:on_adapter_created(domain, adapter)
    if adapter ~= "postgres" then
        return
    end

    print("Adding 'postgres' adapter template to " .. domain)

    local path = appy.get_adapter_root(domain, adapter) .. "adapter.go"
    appy.apply_template({
        template = self.script_root .. "templates/adapter.go",
        target = path,
        args = {
            Module = appy.config.module,
            DomainName = domain
        }
    })

    local path = appy.get_adapter_root(domain, adapter) .. "get.go"
    appy.apply_template({
        template = self.script_root .. "templates/get.go",
        target = path,
        args = {
            Module = appy.config.module,
            DomainName = domain,
            ProjectName = appy.config.project
        }
    })
end

return plugin
