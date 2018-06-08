package main

import (
	"log"
	"github.com/red010b37/ntp-sync-test/ntp"
	"fmt"
)

func main () {

	log.Println("hi")
	pools := []string{"de.pool.ntp.org",
		"jp.pool.ntp.org", "id.pool.ntp.org", "in.pool.ntp.org",
		"0.de.pool.ntp.org", "1.de.pool.ntp.org", "2.de.pool.ntp.org",
		.

		"3.de.ntp.org", "cn.pool.ntp.org", "time.google.com",
		"ntp0.fau.de", "ntps1-0.uni-erlangen.de", "ntps1-0.cs.tu-berlin.de",
		"ptbtime1.ptb.de", "rustime01.rus.uni-stuttgart.de",
		"time1.google.com", "time2.google.com", "time3.google.com", "time4.google.com",
		"hora.rediris.es", "cuco.rediris.es", "hora.roa.es", "ntp.i2t.ehu.es",
		"at.pool.ntp.org", "bg.pool.ntp.org", "ch.pool.ntp.org", "cz.pool.ntp.org",
		"dk.pool.ntp.org", "fi.pool.ntp.org", "fr.pool.ntp.org", "hu.pool.ntp.org",
		"nl.pool.ntp.org", "no.pool.ntp.org", "pl.pool.ntp.org", "ru.pool.ntp.org",
		"uk.pool.ntp.org", "ca.pool.ntp.org", "us.pool.ntp.org", "au.pool.ntp.org",
		"nz.pool.ntp.org", "br.pool.ntp.org"}

	for _, pool := range pools {

		t, err := ntpclient.GetNetworkTime(pool, 123)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(pool, t)

	}

}
