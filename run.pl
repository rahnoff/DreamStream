#!/usr/bin/perl

use strict;
use warnings;

sub run() {
    my $GIT_DIRECTORY = `git rev-parse --show-toplevel`;
    system('docker', 'compose', '-f', "$GIT_REPOSITORY/databases/postgresql/)
}

run()