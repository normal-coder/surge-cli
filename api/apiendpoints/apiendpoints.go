package apiendpoints

var apiPrefix = "/v1"

type featuresAPI struct{ Capture, EnhancedMode, MITM, Rewrite, Scripting, SystemProxy string }
type proxyPolicyAPI struct{ PolicyList, Detail, Test, PolicyGroupList, PolicyGroupsTestResults, PolicyGroupsSelect, PolicyGroupsTest string }
type requestAPI struct{ Recent, Active, Kill string }
type profilesAPI struct{ Current, Reload, Switch, Profiles, Check string }
type dnsAPI struct{ Flush, Dns, DnsDelay string }
type scriptingAPI struct{ Scripting, Evaluate, CronEvaluate string }
type deviceManagementAPI struct{ Devices, Icon string }
type miscAPI struct{ Stop, Events, Rules, Traffic, LogLevel, Ca string }

type ModulesAPI struct{ Current, Events, Modules, Outbound, Traffic string }

var FeaturesToggle = featuresAPI{ // Toggle capabilities
	MITM:         apiPrefix + "/features/mitm",
	Capture:      apiPrefix + "/features/capture",
	Rewrite:      apiPrefix + "/features/rewrite",
	Scripting:    apiPrefix + "/features/scripting",
	SystemProxy:  apiPrefix + "/features/system_proxy",  // Surge Mac Only
	EnhancedMode: apiPrefix + "/features/enhanced_mode", // Surge Mac Only
}

var OutboundMode = apiPrefix + "/outbound" // Use GET to obtain the outbound mode, and use POST to change it.

var ProxyPolicy = proxyPolicyAPI{ // Proxy Policy
	PolicyList:              apiPrefix + "/policies",                   // List all policies.
	Detail:                  apiPrefix + "/policies/detail",            // Obtain the detail of policy.
	Test:                    apiPrefix + "/policies/test",              // Test policies with a URL.
	PolicyGroupList:         apiPrefix + "/policy_groups",              // List all policy groups and their options.
	PolicyGroupsTestResults: apiPrefix + "/policy_groups/test_results", // Obtain the test result of an url-test/fallback/load-balance group.
	PolicyGroupsSelect:      apiPrefix + "/policy_groups/select",       // Obtain the option of a select group.
	PolicyGroupsTest:        apiPrefix + "/policy_groups/test",         // Test a group immediately.
}

var Requests = requestAPI{
	Recent: apiPrefix + "/requests/recent", //List recent requests.
	Active: apiPrefix + "/requests/active", //List all active requests.
	Kill:   apiPrefix + "/requests/kill",   //Kill an active request.

}

var Profiles = profilesAPI{
	Current:  apiPrefix + "/v1/profiles/current?sensitive=0", // Obtain the text content of the current profile. If 'sensitive' is false, all passwords field will be masked.
	Reload:   apiPrefix + "/v1/profiles/reload",              // Execute profile reloading immediately.
	Switch:   apiPrefix + "/v1/profiles/switch",              // Surge Mac Only | Switch to another profile.
	Profiles: apiPrefix + "/v1/profiles",                     // Surge Mac Only 4.0.6+ | Get all available profile names.
	Check:    apiPrefix + "/v1/profiles/check",               // Surge Mac Only 4.0.6+ | Check the profile. If the profile is invalid an error will be returned. Otherwise, the "error" field will be null.
}

var Dns = dnsAPI{
	Flush:    apiPrefix + "/v1/dns/flush",      // Flush the DNS cache.
	Dns:      apiPrefix + "/v1/dns",            // Obtain the current DNS cache content.
	DnsDelay: apiPrefix + "/v1/test/dns_delay", // Test the DNS delay.
}

var Modules = apiPrefix + "/modules" // List the available and enabled modules. / Enable or disable modules.

var Scripting = scriptingAPI{
	Scripting:    apiPrefix + "/scripting",               // List all the scripts.
	Evaluate:     apiPrefix + "/scripting/evaluate",      // Evaluate a script with a mock environment.
	CronEvaluate: apiPrefix + "/scripting/cron/evaluate", // Evaluate a cron script immediately.
}

var DeviceManagement = deviceManagementAPI{
	Devices: apiPrefix + "/devices",      // GET: Obtain the list of the current active and saved devices. | POST: Change the device properties. physicalAddress field is required. And you may adjust multiple or one property from name, address, and shouldHandledBySurge.
	Icon:    apiPrefix + "/devices/icon", // Obtain the icon of a device. You may get the iconID from device.dhcpDevice.icon.(?id={iconID})
}

var Misc = miscAPI{
	Stop:     apiPrefix + "/stop",      // Shutdown Surge engine. If Always On is enabled on Surge iOS, the Surge engine will restart.
	Events:   apiPrefix + "/events",    // Obtain the content of the event center.
	Rules:    apiPrefix + "/rules",     // Obtain the list of rules.
	Traffic:  apiPrefix + "/traffic",   // Obtain traffic information.
	LogLevel: apiPrefix + "/log/level", // Change the log level for the current session.
	Ca:       apiPrefix + "/mitm/ca",   // Obtain the CA certificate for MITM, in DER binary format. (Certificate only, no private key included)
}
