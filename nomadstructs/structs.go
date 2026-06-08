// Specs used for autocompletion
//
// We cannot use the upstream config schemas as they are not publicly exposed.
// Validation uses upstream Nomad parser tho.
package nomadstructs

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/zclconf/go-cty/cty"
)

var JobSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"job",
	},
	TypeName: "job",
	Nested: &hcldec.ObjectSpec{
		"all_at_once": &hcldec.AttrSpec{
			Name: "all_at_once",
			Type: cty.Bool,
		},
		"constraint":    ConstraintSpec,
		"group":         GroupSpec,
		"affinity":      AffinitySpec,
		"vault":         VaultSpec,
		"meta":          MetaSpec,
		"migrate":       MigrateSpec,
		"parameterized": ParameterizedSpec,
		"periodic":      PeriodicSpec,
		"reschedule":    RescheduleSpec,
		"spread":        SpreadSpec,
		"update":        UpdateSpec,

		"datacenters": &hcldec.AttrSpec{
			Name:     "datacenters",
			Type:     cty.List(cty.String),
			Required: true,
		},
		"namespace": &hcldec.AttrSpec{
			Name: "namespace",
			Type: cty.String,
		},
		"priority": &hcldec.AttrSpec{
			Name: "priority",
			Type: cty.Number,
		},
		"region": &hcldec.AttrSpec{
			Name: "region",
			Type: cty.String,
		},
		"type": &hcldec.AttrSpec{
			Name: "type",
			Type: cty.String,
		},
		"vault_token": &hcldec.AttrSpec{
			Name: "vault_token",
			Type: cty.String,
		},
	},
}

var ConstraintSpec = &hcldec.BlockSpec{
	TypeName: "constraint",
	Nested: &hcldec.ObjectSpec{
		"attribute": &hcldec.AttrSpec{
			Name: "attribute",
			Type: cty.String,
		},
		"operator": &hcldec.AttrSpec{
			Name: "operator",
			Type: cty.String,
		},
		"value": &hcldec.AttrSpec{
			Name: "value",
			Type: cty.String,
		},
	},
}

var GroupSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"group",
	},
	TypeName: "group",
	Nested: &hcldec.ObjectSpec{
		"count": &hcldec.AttrSpec{
			Name: "count",
			Type: cty.Number,
		},
		"constraint":     ConstraintSpec,
		"task":           TaskSpec,
		"affinity":       AffinitySpec,
		"vault":          VaultSpec,
		"ephemeral_disk": EphemeralDiskSpec,
		"meta":           MetaSpec,
		"migrate":        MigrateSpec,
		"reschedule":     RescheduleSpec,
		"restart":        RestartSpec,
		"spread":         SpreadSpec,
		"update":         UpdateSpec,
		"network":        NetworkSpec,
		"service":        ServiceSpec,
		"volume":         VolumeSpec,
	},
}

var VolumeMountSpec = &hcldec.BlockSpec{
	TypeName: "config",
	Nested: &hcldec.ObjectSpec{
		"volume": &hcldec.AttrSpec{
			Name: "volume",
			Type: cty.String,
		},
		"destination": &hcldec.AttrSpec{
			Name: "destination",
			Type: cty.String,
		},
		"read_only": &hcldec.AttrSpec{
			Name: "read_only",
			Type: cty.Bool,
		},
		"propagation_mode": &hcldec.AttrSpec{
			Name: "propagation_mode",
			Type: cty.String,
		},
		"selinux_label": &hcldec.AttrSpec{
			Name: "selinux_label",
			Type: cty.String,
		},
	},
}
var LifecycleSpec = &hcldec.BlockSpec{
	TypeName: "lifecycle",
	Nested: &hcldec.ObjectSpec{
		"hook": &hcldec.AttrSpec{
			Name: "hook",
			Type: cty.String,
		},
		"sidecar": &hcldec.AttrSpec{
			Name: "sidecar",
			Type: cty.Bool,
		},
	},
}

var ActionSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"action",
	},
	TypeName: "action",
	Nested: &hcldec.ObjectSpec{
		"command": &hcldec.AttrSpec{
			Name: "command",
			Type: cty.String,
		},
		"args": &hcldec.AttrSpec{
			Name: "args",
			Type: cty.List(cty.String),
		},
	},
}

var TaskSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"task",
	},
	TypeName: "task",
	Nested: &hcldec.ObjectSpec{
		"constraint":       ConstraintSpec,
		"artifact":         ArtifactSpec,
		"affinity":         AffinitySpec,
		"vault":            VaultSpec,
		"dispatch_payload": DispatchPayloadSpec,
		"env":              EnvSpec,
		"logs":             LogsSpec,
		"meta":             MetaSpec,
		"resources":        ResourcesSpec,
		"service":          ServiceSpec,
		"restart":          RestartSpec,
		"template":         TemplateSpec,
		"volume_mount":     VolumeMountSpec,
		"lifecycle":        LifecycleSpec,
		"action":           ActionSpec,

		"driver": &hcldec.AttrSpec{
			Name:     "driver",
			Type:     cty.String,
			Required: true,
		},
		"kill_timeout": &hcldec.AttrSpec{
			Name: "kill_timeout",
			Type: cty.String,
		},
		"kill_signal": &hcldec.AttrSpec{
			Name: "kill_signal",
			Type: cty.String,
		},
		"leader": &hcldec.AttrSpec{
			Name: "leader",
			Type: cty.Bool,
		},
		"shutdown_delay": &hcldec.AttrSpec{
			Name: "shutdown_delay",
			Type: cty.String,
		},
		"user": &hcldec.AttrSpec{
			Name: "user",
			Type: cty.String,
		},

		"config": &hcldec.BlockAttrsSpec{
			TypeName:    "config",
			ElementType: cty.DynamicPseudoType,
		},
	},
}

var ArtifactSpec = &hcldec.BlockSpec{
	TypeName: "artifact",
	Nested: &hcldec.ObjectSpec{
		"destination": &hcldec.AttrSpec{
			Name: "destination",
			Type: cty.String,
		},
		"mode": &hcldec.AttrSpec{
			Name: "mode",
			Type: cty.String,
		},
		"options": &hcldec.BlockAttrsSpec{
			TypeName:    "options",
			ElementType: cty.String,
		},
		"source": &hcldec.AttrSpec{
			Name:     "source",
			Type:     cty.String,
			Required: true,
		},
	},
}

var AffinitySpec = &hcldec.BlockSpec{
	TypeName: "affinity",
	Nested: &hcldec.ObjectSpec{
		"attribute": &hcldec.AttrSpec{
			Name: "attribute",
			Type: cty.String,
		},
		"operator": &hcldec.AttrSpec{
			Name: "operator",
			Type: cty.String,
		},
		"value": &hcldec.AttrSpec{
			Name: "value",
			Type: cty.String,
		},
		"weight": &hcldec.AttrSpec{
			Name: "weight",
			Type: cty.String,
		},
	},
}

var VaultSpec = &hcldec.BlockSpec{
	TypeName: "vault",
	Nested: &hcldec.ObjectSpec{
		"change_mode": &hcldec.AttrSpec{
			Name: "change_mode",
			Type: cty.String,
		},
		"change_signal": &hcldec.AttrSpec{
			Name: "change_signal",
			Type: cty.String,
		},
		"env": &hcldec.AttrSpec{
			Name: "env",
			Type: cty.Bool,
		},
		"policies": &hcldec.AttrSpec{
			Name: "policies",
			Type: cty.Set(cty.String),
		},
	},
}

var DeviceSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"device",
	},
	TypeName: "device",
	Nested: &hcldec.ObjectSpec{
		"name": &hcldec.AttrSpec{
			Name: "name",
			Type: cty.String,
		},
		"count": &hcldec.AttrSpec{
			Name: "count",
			Type: cty.Number,
		},
		"constraint": ConstraintSpec,
		"affinity":   AffinitySpec,
	},
}

var CheckRestartSpec = &hcldec.BlockSpec{
	TypeName: "check_restart",
	Nested: &hcldec.ObjectSpec{
		"limit": &hcldec.AttrSpec{
			Name: "limit",
			Type: cty.Number,
		},
		"grace": &hcldec.AttrSpec{
			Name: "grace",
			Type: cty.String,
		},
		"ignore_warnings": &hcldec.AttrSpec{
			Name: "ignore_warnings",
			Type: cty.Bool,
		},
	},
}

var DispatchPayloadSpec = &hcldec.BlockSpec{
	TypeName: "dispatch_payload",
	Nested: &hcldec.ObjectSpec{
		"file": &hcldec.AttrSpec{
			Name: "file",
			Type: cty.String,
		},
	},
}

var EnvSpec = &hcldec.BlockAttrsSpec{
	TypeName:    "env",
	ElementType: cty.Map(cty.String),
}

var EphemeralDiskSpec = &hcldec.BlockSpec{
	TypeName: "ephemeral_disk",
	Nested: &hcldec.ObjectSpec{
		"migrate": &hcldec.AttrSpec{
			Name: "migrate",
			Type: cty.Bool,
		},
		"size": &hcldec.AttrSpec{
			Name: "size",
			Type: cty.Number,
		},
		"sticky": &hcldec.AttrSpec{
			Name: "sticky",
			Type: cty.Bool,
		},
	},
}

var LogsSpec = &hcldec.BlockSpec{
	TypeName: "logs",
	Nested: &hcldec.ObjectSpec{
		"max_files": &hcldec.AttrSpec{
			Name: "max_files",
			Type: cty.Number,
		},
		"max_file_size": &hcldec.AttrSpec{
			Name: "max_file_size",
			Type: cty.Number,
		},
	},
}

var MetaSpec = &hcldec.BlockAttrsSpec{
	TypeName:    "meta",
	ElementType: cty.String,
}

var MigrateSpec = &hcldec.BlockSpec{
	TypeName: "migrate",
	Nested: &hcldec.ObjectSpec{
		"max_parallel": &hcldec.AttrSpec{
			Name: "max_parallel",
			Type: cty.Number,
		},
		"health_check": &hcldec.AttrSpec{
			Name: "health_check",
			Type: cty.String,
		},
		"min_healthy_time": &hcldec.AttrSpec{
			Name: "min_healthy_time",
			Type: cty.String,
		},
		"healthy_deadline": &hcldec.AttrSpec{
			Name: "healthy_deadline",
			Type: cty.String,
		},
	},
}

var DNSSpec = &hcldec.BlockSpec{
	TypeName: "dns",
	Nested: &hcldec.ObjectSpec{
		"servers": &hcldec.AttrSpec{
			Name: "servers",
			Type: cty.List(cty.String),
		},
		"searches": &hcldec.AttrSpec{
			Name: "searches",
			Type: cty.List(cty.String),
		},
		"options": &hcldec.AttrSpec{
			Name: "options",
			Type: cty.List(cty.String),
		},
	},
}
var CNISpec = &hcldec.BlockSpec{
	TypeName: "cni",
	Nested: &hcldec.ObjectSpec{
		"args": &hcldec.AttrSpec{
			Name: "args",
			Type: cty.Map(cty.String),
		},
	},
}
var PortSpec = &hcldec.BlockMapSpec{
	LabelNames: []string{
		"port",
	},
	TypeName: "port",
	Nested: &hcldec.ObjectSpec{
		"static": &hcldec.AttrSpec{
			Name: "static",
			Type: cty.Number,
		},
		"to": &hcldec.AttrSpec{
			Name: "to",
			Type: cty.Number,
		},
		"host_network": &hcldec.AttrSpec{
			Name: "host_network",
			Type: cty.String,
		},
		"ignore_collision": &hcldec.AttrSpec{
			Name: "ignore_collision",
			Type: cty.Bool,
		},
	},
}
var NetworkSpec = &hcldec.BlockSpec{
	TypeName: "network",
	Nested: &hcldec.ObjectSpec{
		"dns":  DNSSpec,
		"cni":  CNISpec,
		"port": PortSpec,

		"mode": &hcldec.AttrSpec{
			Name: "mode",
			Type: cty.String,
		},
		"hostname": &hcldec.AttrSpec{
			Name: "hostname",
			Type: cty.String,
		},
	},
}

var ParameterizedSpec = &hcldec.BlockSpec{
	TypeName: "parameterized",
	Nested: &hcldec.ObjectSpec{
		"meta_optional": &hcldec.AttrSpec{
			Name: "meta_optional",
			Type: cty.List(cty.String),
		},
		"meta_required": &hcldec.AttrSpec{
			Name: "meta_required",
			Type: cty.List(cty.String),
		},
		"payload": &hcldec.AttrSpec{
			Name: "payload",
			Type: cty.String,
		},
	},
}

var PeriodicSpec = &hcldec.BlockSpec{
	TypeName: "periodic",
	Nested: &hcldec.ObjectSpec{
		"cron": &hcldec.AttrSpec{
			Name: "cron",
			Type: cty.String,
		},
		"crons": &hcldec.AttrSpec{
			Name: "crons",
			Type: cty.List(cty.String),
		},
		"prohibit_overlap": &hcldec.AttrSpec{
			Name: "prohibit_overlap",
			Type: cty.Bool,
		},
		"time_zone": &hcldec.AttrSpec{
			Name: "time_zone",
			Type: cty.String,
		},
		"enabled": &hcldec.AttrSpec{
			Name: "enabled",
			Type: cty.Bool,
		},
	},
}

var RescheduleSpec = &hcldec.BlockSpec{
	TypeName: "reschedule",
	Nested: &hcldec.ObjectSpec{
		"attempts": &hcldec.AttrSpec{
			Name: "attempts",
			Type: cty.Number,
		},
		"interval": &hcldec.AttrSpec{
			Name: "interval",
			Type: cty.String,
		},
		"delay": &hcldec.AttrSpec{
			Name: "delay",
			Type: cty.String,
		},
		"delay_function": &hcldec.AttrSpec{
			Name: "delay_function",
			Type: cty.String,
		},
		"max_delay": &hcldec.AttrSpec{
			Name: "max_delay",
			Type: cty.String,
		},
		"unlimited": &hcldec.AttrSpec{
			Name: "unlimited",
			Type: cty.Bool,
		},
	},
}

var ResourcesSpec = &hcldec.BlockSpec{
	TypeName: "resources",
	Nested: &hcldec.ObjectSpec{
		"cpu": &hcldec.AttrSpec{
			Name: "cpu",
			Type: cty.Number,
		},
		"memory": &hcldec.AttrSpec{
			Name: "memory",
			Type: cty.Number,
		},
		"network": NetworkSpec,
		"device":  DeviceSpec,
	},
}

var RestartSpec = &hcldec.BlockSpec{
	TypeName: "restart",
	Nested: &hcldec.ObjectSpec{
		"attempts": &hcldec.AttrSpec{
			Name: "attempts",
			Type: cty.Number,
		},
		"delay": &hcldec.AttrSpec{
			Name: "delay",
			Type: cty.String,
		},
		"interval": &hcldec.AttrSpec{
			Name: "interval",
			Type: cty.String,
		},
		"mode": &hcldec.AttrSpec{
			Name: "mode",
			Type: cty.String,
		},
		"render_templates": &hcldec.AttrSpec{
			Name: "render_templates",
			Type: cty.Bool,
		},
	},
}

var ServiceSpec = &hcldec.BlockSpec{
	TypeName: "service",
	Nested: &hcldec.ObjectSpec{
		"name": &hcldec.AttrSpec{
			Name: "name",
			Type: cty.String,
		},
		"port": &hcldec.AttrSpec{
			Name: "port",
			Type: cty.String,
		},
		"provider": &hcldec.AttrSpec{
			Name: "provider",
			Type: cty.String,
		},
		"tags": &hcldec.AttrSpec{
			Name: "tags",
			Type: cty.List(cty.String),
		},
		"canary_tags": &hcldec.AttrSpec{
			Name: "canary_tags",
			Type: cty.List(cty.String),
		},
		"address_mode": &hcldec.AttrSpec{
			Name: "address_mode",
			Type: cty.String,
		},
		"check": &hcldec.BlockSetSpec{
			TypeName: "check",
			Nested: &hcldec.ObjectSpec{
				"address_mode": &hcldec.AttrSpec{
					Name: "address_mode",
					Type: cty.String,
				},
				"args": &hcldec.AttrSpec{
					Name: "args",
					Type: cty.List(cty.String),
				},
				"check_restart": CheckRestartSpec,
				"command": &hcldec.AttrSpec{
					Name: "command",
					Type: cty.String,
				},
				"grpc_service": &hcldec.AttrSpec{
					Name: "grpc_service",
					Type: cty.String,
				},
				"grpc_use_tls": &hcldec.AttrSpec{
					Name: "grpc_use_tls",
					Type: cty.Bool,
				},
				"initial_status": &hcldec.AttrSpec{
					Name: "initial_status",
					Type: cty.String,
				},
				"interval": &hcldec.AttrSpec{
					Name: "interval",
					Type: cty.String,
				},
				"method": &hcldec.AttrSpec{
					Name: "method",
					Type: cty.String,
				},
				"name": &hcldec.AttrSpec{
					Name: "name",
					Type: cty.String,
				},
				"path": &hcldec.AttrSpec{
					Name: "path",
					Type: cty.String,
				},
				"port": &hcldec.AttrSpec{
					Name: "port",
					Type: cty.String,
				},
				"protocol": &hcldec.AttrSpec{
					Name: "protocol",
					Type: cty.String,
				},
				"timeout": &hcldec.AttrSpec{
					Name:     "timeout",
					Type:     cty.String,
					Required: true,
				},
				"type": &hcldec.AttrSpec{
					Name:     "type",
					Type:     cty.String,
					Required: true,
				},
				"tls_skip_verify": &hcldec.AttrSpec{
					Name: "tls_skip_verify",
					Type: cty.Bool,
				},
				"header": &hcldec.BlockAttrsSpec{
					TypeName:    "header",
					ElementType: cty.String,
				},
			},
		},
	},
}

var SpreadSpec = &hcldec.BlockSpec{
	TypeName: "spread",
	Nested: &hcldec.ObjectSpec{
		"attribute": &hcldec.AttrSpec{
			Name: "attribute",
			Type: cty.String,
		},
		"weight": &hcldec.AttrSpec{
			Name: "weight",
			Type: cty.Number,
		},
		"target": &hcldec.BlockMapSpec{
			LabelNames: []string{
				"target",
			},
			TypeName: "target",
			Nested: &hcldec.ObjectSpec{
				"value": &hcldec.AttrSpec{
					Name: "value",
					Type: cty.String,
				},
				"percent": &hcldec.AttrSpec{
					Name: "percent",
					Type: cty.String,
				},
			},
		},
	},
}

var TemplateSpec = &hcldec.BlockSpec{
	TypeName: "template",
	Nested: &hcldec.ObjectSpec{
		"change_mode": &hcldec.AttrSpec{
			Name: "change_mode",
			Type: cty.String,
		},
		"change_signal": &hcldec.AttrSpec{
			Name: "change_signal",
			Type: cty.String,
		},
		"data": &hcldec.AttrSpec{
			Name: "data",
			Type: cty.String,
		},
		"destination": &hcldec.AttrSpec{
			Name:     "destination",
			Type:     cty.String,
			Required: true,
		},
		"env": &hcldec.AttrSpec{
			Name: "env",
			Type: cty.Bool,
		},
		"error_on_missing_key": &hcldec.AttrSpec{
			Name: "error_on_missing_key",
			Type: cty.Bool,
		},
		"left_delimiter": &hcldec.AttrSpec{
			Name: "left_delimiter",
			Type: cty.String,
		},
		"perms": &hcldec.AttrSpec{
			Name: "perms",
			Type: cty.String,
		},
		"uid": &hcldec.AttrSpec{
			Name: "uid",
			Type: cty.Number,
		},
		"gid": &hcldec.AttrSpec{
			Name: "gid",
			Type: cty.Number,
		},
		"once": &hcldec.AttrSpec{
			Name: "once",
			Type: cty.Bool,
		},
		"right_delimiter": &hcldec.AttrSpec{
			Name: "right_delimiter",
			Type: cty.String,
		},
		"source": &hcldec.AttrSpec{
			Name: "source",
			Type: cty.String,
		},
		"splay": &hcldec.AttrSpec{
			Name: "splay",
			Type: cty.String,
		},
		"vault_grace": &hcldec.AttrSpec{
			Name: "vault_grace",
			Type: cty.String,
		},
	},
}

var UpdateSpec = &hcldec.BlockSpec{
	TypeName: "update",
	Nested: &hcldec.ObjectSpec{
		"max_parallel": &hcldec.AttrSpec{
			Name: "max_parallel",
			Type: cty.Number,
		},
		"health_check": &hcldec.AttrSpec{
			Name: "health_check",
			Type: cty.String,
		},
		"min_healthy_time": &hcldec.AttrSpec{
			Name: "min_healthy_time",
			Type: cty.String,
		},
		"healthy_deadline": &hcldec.AttrSpec{
			Name: "healthy_deadline",
			Type: cty.String,
		},
		"progress_deadline": &hcldec.AttrSpec{
			Name: "progress_deadline",
			Type: cty.String,
		},
		"auto_revert": &hcldec.AttrSpec{
			Name: "auto_revert",
			Type: cty.Bool,
		},
		"canary": &hcldec.AttrSpec{
			Name: "canary",
			Type: cty.Number,
		},
		"stagger": &hcldec.AttrSpec{
			Name: "stagger",
			Type: cty.String,
		},
	},
}

var VolumeSpec = &hcldec.BlockMapSpec{
	TypeName: "volume",
	LabelNames: []string{
		"volume",
	},
	Nested: &hcldec.ObjectSpec{
		"type": &hcldec.AttrSpec{
			Name: "type",
			Type: cty.String,
		},
		"source": &hcldec.AttrSpec{
			Name:     "source",
			Type:     cty.String,
			Required: true,
		},
		"read_only": &hcldec.AttrSpec{
			Name: "read_only",
			Type: cty.Bool,
		},
		"sticky": &hcldec.AttrSpec{
			Name: "sticky",
			Type: cty.Bool,
		},
		"per_alloc": &hcldec.AttrSpec{
			Name: "per_alloc",
			Type: cty.Bool,
		},
		"access_mode": &hcldec.AttrSpec{
			Name: "access_mode",
			Type: cty.String,
		},
		"attachment_mode": &hcldec.AttrSpec{
			Name: "attachment_mode",
			Type: cty.String,
		},
		"mount_options": &hcldec.BlockSpec{
			TypeName: "mount_options",
			Nested: &hcldec.ObjectSpec{
				"fs_type": &hcldec.AttrSpec{
					Name: "fs_type",
					Type: cty.String,
				},
				"mount_flags": &hcldec.AttrSpec{
					Name: "attachment_mode",
					Type: cty.List(cty.String),
				},
			},
		},
	},
}
