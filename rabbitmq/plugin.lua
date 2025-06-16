local appy = require("appy")

local plugin = {}

function plugin:on_adapter_created(domain, adapter)
    if adapter ~= "rabbitmq" then
        return
    end

    print("Adding 'rabbitmq' adapter template to " .. domain)
    appy.apply_template({
        template = self.script_root .. "templates/adapter.go",
        target = appy.get_adapter_root(domain, adapter) .. "adapter.go",
        args = {
            Module = appy.config.module,
            DomainName = domain
        }
    })

    appy.apply_template({
        template = self.script_root .. "templates/send.go",
        target = appy.get_adapter_root(domain, adapter) .. "send.go",
        args = {
            Module = appy.config.module,
            DomainName = domain
        }
    })

    appy.apply_template({
        template = self.script_root .. "templates/receive.go",
        target = appy.get_adapter_root(domain, adapter) .. "receive.go",
        args = {
            Module = appy.config.module,
            DomainName = domain
        }
    })
end

return plugin
