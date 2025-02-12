package taglib

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Parser", func() {
	var e *Parser
	BeforeEach(func() {
		e = &Parser{}
	})
	Context("Parse", func() {
		It("correctly parses metadata from all files in folder", func() {
			mds, err := e.Parse("tests/fixtures/test.mp3", "tests/fixtures/test.ogg")
			Expect(err).NotTo(HaveOccurred())
			Expect(mds).To(HaveLen(2))

			m := mds["tests/fixtures/test.mp3"]
			Expect(m).To(HaveKeyWithValue("title", []string{"Song", "Song"}))
			Expect(m).To(HaveKeyWithValue("album", []string{"Album", "Album"}))
			Expect(m).To(HaveKeyWithValue("artist", []string{"Artist", "Artist"}))
			Expect(m).To(HaveKeyWithValue("albumartist", []string{"Album Artist"}))
			Expect(m).To(HaveKeyWithValue("tcmp", []string{"1"})) // Compilation
			Expect(m).To(HaveKeyWithValue("genre", []string{"Rock"}))
			Expect(m).To(HaveKeyWithValue("date", []string{"2014", "2014"}))
			Expect(m).To(HaveKeyWithValue("tracknumber", []string{"2/10", "2/10", "2"}))
			Expect(m).To(HaveKeyWithValue("discnumber", []string{"1/2"}))
			Expect(m).To(HaveKeyWithValue("has_picture", []string{"true"}))
			Expect(m).To(HaveKeyWithValue("duration", []string{"1.02"}))
			Expect(m).To(HaveKeyWithValue("bitrate", []string{"192"}))
			Expect(m).To(HaveKeyWithValue("comment", []string{"Comment1\nComment2"}))
			Expect(m).To(HaveKeyWithValue("lyrics", []string{"Lyrics 1\rLyrics 2"}))
			Expect(m).To(HaveKeyWithValue("bpm", []string{"123"}))

			m = mds["tests/fixtures/test.ogg"]
			Expect(err).To(BeNil())
			Expect(m).ToNot(HaveKey("title"))
			Expect(m).ToNot(HaveKey("has_picture"))
			Expect(m).To(HaveKeyWithValue("duration", []string{"1.04"}))
			Expect(m).To(HaveKeyWithValue("fbpm", []string{"141.7"}))

			// TabLib 1.12 returns 18, previous versions return 39.
			// See https://github.com/taglib/taglib/commit/2f238921824741b2cfe6fbfbfc9701d9827ab06b
			Expect(m).To(HaveKey("bitrate"))
			Expect(m["bitrate"][0]).To(BeElementOf("18", "39"))
		})
	})
})
