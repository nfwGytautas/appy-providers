local appy = require("appy")

local plugin = {}

function plugin:on_adapter_created(domain, adapter)
    if adapter ~= "http" then
        return
    end

    print("Adding 'http' adapter template to " .. domain)
    local path = appy.get_adapter_root(domain, adapter) .. "http.go"

    appy.apply_template({
        template = self.script_root .. "templates/adapter.go",
        target = path,
        args = {
            DomainName = domain
        }
    })
end

return plugin
