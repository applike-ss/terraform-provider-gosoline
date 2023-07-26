package builder

func NewPanelStreamProducerDaemonSizes(producer MetadataStreamProducer) PanelFactory {
	return func(resourceNames ResourceNames, gridPos PanelGridPos) Panel {
		return Panel{
			Datasource: resourceNames.GrafanaCloudWatchDatasourceName,
			FieldConfig: PanelFieldConfig{
				Defaults: PanelFieldConfigDefaults{
					Custom: PanelFieldConfigDefaultsCustom{
						LineWidth:     2,
						AxisPlacement: "right",
					},
					Min: "0",
				},
				Overrides: []PanelFieldConfigOverwrite{},
			},
			GridPos: gridPos,
			Targets: []interface{}{
				PanelTargetCloudWatch{
					Alias: "Batch Size",
					Dimensions: map[string]string{
						"ProducerDaemon": producer.Name,
					},
					MatchExact: false,
					MetricName: "BatchSize",
					Namespace:  resourceNames.CloudwatchNamespace,
					Period:     "",
					RefId:      "A",
					Region:     "default",
					Statistics: []string{
						"Average",
					},
				}, PanelTargetCloudWatch{
					Alias: "Aggregate Size",
					Dimensions: map[string]string{
						"ProducerDaemon": producer.Name,
					},
					MatchExact: false,
					MetricName: "AggregateSize",
					Namespace:  resourceNames.CloudwatchNamespace,
					Period:     "",
					RefId:      "B",
					Region:     "default",
					Statistics: []string{
						"Average",
					},
				},
			},
			Options: &PanelOptionsCloudWatch{},
			Title:   "Average Batch Size / Aggregation Size",
			Type:    "timeseries",
		}
	}
}

func NewPanelStreamProducerMessageCount(producer MetadataStreamProducer) PanelFactory {
	return func(resourceNames ResourceNames, gridPos PanelGridPos) Panel {
		return Panel{
			Datasource: resourceNames.GrafanaCloudWatchDatasourceName,
			FieldConfig: PanelFieldConfig{
				Defaults: PanelFieldConfigDefaults{
					Custom: PanelFieldConfigDefaultsCustom{
						LineWidth:     2,
						AxisPlacement: "right",
					},
					Min: "0",
				},
				Overrides: []PanelFieldConfigOverwrite{},
			},
			GridPos: gridPos,
			Targets: []interface{}{
				PanelTargetCloudWatch{
					Alias: "Message Count",
					Dimensions: map[string]string{
						"ProducerDaemon": producer.Name,
					},
					MatchExact: false,
					MetricName: "MessageCount",
					Namespace:  resourceNames.CloudwatchNamespace,
					Period:     "",
					RefId:      "A",
					Region:     "default",
					Statistics: []string{
						"Sum",
					},
				},
			},
			Options: &PanelOptionsCloudWatch{},
			Title:   "Message Count",
			Type:    "timeseries",
		}
	}
}
