package audio

var (
	// MONO

	// FormatMono22050 is mono 22.05kHz format.
	FormatMono22050 = &Format{
		NumChannels: 1,
		SampleRate:  22050,
	}
	// FormatMono44100 is mono 8bit 44.1kHz format.
	FormatMono44100 = &Format{
		NumChannels: 1,
		SampleRate:  44100,
	}
	// FormatMono48000 is mono 48kHz format.
	FormatMono48000 = &Format{
		NumChannels: 1,
		SampleRate:  48000,
	}
	// FormatMono96000 is mono 96kHz format.
	FormatMono96000 = &Format{
		NumChannels: 1,
		SampleRate:  96000,
	}

	// STEREO

	// FormatStereo22050 is stereo 22.05kHz format.
	FormatStereo22050 = &Format{
		NumChannels: 2,
		SampleRate:  22050,
	}
	// FormatStereo44100 is stereo 8bit 44.1kHz format.
	FormatStereo44100 = &Format{
		NumChannels: 2,
		SampleRate:  44100,
	}
	// FormatStereo48000 is stereo 48kHz format.
	FormatStereo48000 = &Format{
		NumChannels: 2,
		SampleRate:  48000,
	}
	// FormatStereo96000 is stereo 96kHz format.
	FormatStereo96000 = &Format{
		NumChannels: 2,
		SampleRate:  96000,
	}
)
