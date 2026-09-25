package download

import (
	"testing"
)

func TestParseMediaType(t *testing.T) {
	tests := []struct {
		in      string
		want    MediaType
		wantErr bool
	}{
		{"mp3", TYPE_MP3, false},
		{"MP3", TYPE_MP3, false},
		{" mp4 ", TYPE_MP4, false},
		{"wav", "", true},
		{"", "", true},
	}
	for _, tt := range tests {
		got, err := ParseMediaType(tt.in)
		if (err != nil) != tt.wantErr {
			t.Errorf("ParseMediaType(%q) error = %v, wantErr %v", tt.in, err, tt.wantErr)
			continue
		}
		if got != tt.want {
			t.Errorf("ParseMediaType(%q) = %q, want %q", tt.in, got, tt.want)
		}
	}
}

func TestValidateQuality(t *testing.T) {
	for tier := QUALITY_MIN; tier <= QUALITY_MAX; tier++ {
		if err := ValidateQuality(tier); err != nil {
			t.Errorf("ValidateQuality(%d) = %v, want nil", tier, err)
		}
	}
	if err := ValidateQuality(0); err == nil {
		t.Error("ValidateQuality(0) = nil, want error")
	}
	if err := ValidateQuality(6); err == nil {
		t.Error("ValidateQuality(6) = nil, want error")
	}
}

func TestEveryTierIsMapped(t *testing.T) {
	for tier := QUALITY_MIN; tier <= QUALITY_MAX; tier++ {
		if _, ok := videoHeights[tier]; !ok {
			t.Errorf("videoHeights missing tier %d", tier)
		}
		if _, ok := audioBitrates[tier]; !ok {
			t.Errorf("audioBitrates missing tier %d", tier)
		}
	}
}

func TestTiersAscendMonotonically(t *testing.T) {
	// Video: 0 means uncapped, so skip it for the ascending check.
	prev := 0
	for tier := QUALITY_MIN; tier <= QUALITY_MAX; tier++ {
		h := videoHeights[tier]
		if tier < QUALITY_MAX && h <= prev {
			t.Errorf("videoHeights[%d]=%d not > prev=%d", tier, h, prev)
		}
		if tier < QUALITY_MAX {
			prev = h
		}
	}
}

func TestDescribe(t *testing.T) {
	tests := []struct {
		mt   MediaType
		tier int
		want string
	}{
		{TYPE_MP3, 1, "64K"},
		{TYPE_MP3, 3, "192K"},
		{TYPE_MP3, 5, "320K"},
		{TYPE_MP4, 1, "360p"},
		{TYPE_MP4, 3, "720p"},
		{TYPE_MP4, 5, "best"},
		{TYPE_MP3, 0, "?"},
	}
	for _, tt := range tests {
		got := Describe(tt.mt, tt.tier)
		if got != tt.want {
			t.Errorf("Describe(%s, %d) = %q, want %q", tt.mt, tt.tier, got, tt.want)
		}
	}
}

func TestQualityTable(t *testing.T) {
	table := QualityTable(TYPE_MP4)
	if table == "" {
		t.Error("QualityTable(mp4) returned empty string")
	}
}
