package main

// This test needs to be updated for a "nil" default backend
//
// func TestInitConfigDefaultConfig(t *testing.T) {
// 	log.SetLevel("warn")
// 	want := Config{
// 		BackendsConfig: BackendsConfig{
// 			Backend:      "etcd",
// 			BackendNodes: []string{"http://127.0.0.1:2379"},
// 			Scheme:       "http",
// 			Filter:       "*",
// 		},
// 		TemplateConfig: TemplateConfig{
// 			ConfDir:     "/etc/confd",
// 			ConfigDir:   "/etc/confd/conf.d",
// 			TemplateDir: "/etc/confd/templates",
// 			Noop:        false,
// 		},
// 		ConfigFile: "/etc/confd/confd.toml",
// 		Interval:   600,
// 	}
// 	if err := initConfig(); err != nil {
// 		t.Errorf(err.Error())
// 	}
// 	if !reflect.DeepEqual(want, config) {
// 		t.Errorf("initConfig() = %v, want %v", config, want)
// 	}
// }
