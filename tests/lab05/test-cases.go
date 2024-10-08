blood

{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap10.html", 0.010920563015693252},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap24.html", 0.0095021521968651},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap16.html", 0.0069545059403071565},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap03.html", 0.005205320994794678},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap12.html", 0.004026845637583892},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap11.html", 0.003806140573458513},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap04.html", 0.003499494517458589},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap13.html", 0.003258272391571935},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap22.html", 0.0032243618450515},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap15.html", 0.0030999741668819424},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap09.html", 0.002285037237643872},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap14.html", 0.0021865889212827985},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap18.html", 0.0019421665947345703},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap02.html", 0.0015434745326702106},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap19.html", 0.0014889569029696416},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap25.html", 0.0013512499061632008},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap08.html", 0.0012494793835901706},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap27.html", 0.001105990783410138},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap06.html", 0.0007259820924417197},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap20.html", 0.000710732054015636},
{"http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/chap07.html", 0.0006302521008403361},

func TestCrawlDelay(t *testing.T) {
	t1 := time.Now()
	idx := NewIndex()
	// seems like serve is already running
	// serve(idx)
	crawl(idx, "http://localhost:8080/top10/Dracula%20%7C%20Project%20Gutenberg/index.html")
	t2 := time.Now()
	if t2.Sub(t1) < (10 * time.Second) {
		t.Errorf("TestDisallow was too fast\n")
	}
}
