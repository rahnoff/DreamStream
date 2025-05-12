CREATE KEYSPACE IF NOT EXISTS dream_stream WITH REPLICATION =
{ 
    'class': 'NetworkTopologyStrategy', 
    'datacenter1': 1 
}
AND DURABLE_WRITES = true;
