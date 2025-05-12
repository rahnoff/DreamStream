use strict;
use warnings;

use DBI;


sub customer_emulator()
{
    # Cassandra
    my $cassandra_host1       = 'localhost';
    my $cassandra_host2       = '192.168.0.2';
    my $cassandra_host3       = '192.168.0.3';
    my $cassandra_username    = "ServiceUserName";
    my $cassandra_password    = "ServicePassword";
    my $cassandra_port        = '9142';
    my $cassandra_keyspace    = 'dream_stream';
    my $cassandra_consistency = 'local_quorum';
    my $cassandra_tls         = '0';
    # my $cassandra_dsn         = "dbi:Cassandra:keyspace=$cassandra_keyspace;
    #                              host=$cassandra_host1,$cassandra_host2,$cassandra_host3;
    #                              port=$cassandra_port;
    #                              consistency=$cassandra_consistency;
    #                              tls=$cassandra_tls;";
    my $cassandra_dsn         = "dbi:Cassandra:keyspace=$cassandra_keyspace;
                                 host=$cassandra_host1;
                                 port=$cassandra_port;
                                 consistency=$cassandra_consistency;
                                 tls=$cassandra_tls;";
    my $cassandra_handle      = DBI->connect($cassandra_dsn, $cassandra_username, $cassandra_password);
    $cassandra_handle->do("INSERT INTO courses (id, field_one, field_two) VALUES (?, ?, ?)", { Consistency => "quorum" }, 1, "String value", 38962986);
    $cassandra_handle->disconnect();
}


sub main()
{
    customer_emulator();
}


main();
