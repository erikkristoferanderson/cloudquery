# Table: aws_elasticache_update_actions

This table shows data for Elasticache Update Actions.

https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_UpdateAction.html

The composite primary key for this table is (**account_id**, **region**, **cache_cluster_id**, **replication_group_id**, **service_update_name**).

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|account_id (PK)|`utf8`|
|region (PK)|`utf8`|
|cache_cluster_id (PK)|`utf8`|
|cache_node_update_status|`json`|
|engine|`utf8`|
|estimated_update_time|`utf8`|
|node_group_update_status|`json`|
|nodes_updated|`utf8`|
|replication_group_id (PK)|`utf8`|
|service_update_name (PK)|`utf8`|
|service_update_recommended_apply_by_date|`timestamp[us, tz=UTC]`|
|service_update_release_date|`timestamp[us, tz=UTC]`|
|service_update_severity|`utf8`|
|service_update_status|`utf8`|
|service_update_type|`utf8`|
|sla_met|`utf8`|
|update_action_available_date|`timestamp[us, tz=UTC]`|
|update_action_status|`utf8`|
|update_action_status_modified_date|`timestamp[us, tz=UTC]`|