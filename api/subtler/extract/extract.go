package extract

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"pier/api/subtler/types"
	"slices"
	"strings"

	"github.com/google/uuid"
)

func Extract(video string) (string, error) {
	log := ""
	languages := strings.Split(os.Getenv("SUBTLER_LANGS"), ",")

	infoCmd := os.Getenv("SUBTLER_INFO")
	run := exec.Command(infoCmd, "--identification-format", "json", "--identify", video)
	output, err := run.CombinedOutput()
	if err != nil {
		return log, err
	}

	var info map[string]interface{}
	json.Unmarshal(output, &info)

	var subtitleTracks []types.SubtitleTrack
	tracks := info["tracks"].([]interface{})
	for _, track := range tracks {
		track := track.(map[string]interface{})
		trackType := track["type"].(string)
		if trackType != "subtitles" {
			continue
		}

		properties := track["properties"].(map[string]interface{})
		var trackLang string
		trackLangField := properties["language"]
		if trackLangField != nil {
			trackLang = trackLangField.(string)
		} else {
			trackLang = "und"
		}

		trackLang = strings.ToLower(trackLang)
		if !slices.Contains(languages, trackLang) {
			continue
		}

		var trackName string
		trackNameField := properties["track_name"]
		if trackNameField != nil {
			trackName = trackNameField.(string)
		} else {
			trackName = "Subtitle"
		}

		trackJson, err := json.MarshalIndent(track, "", "  ")
		if err != nil {
			return log, err
		}

		log = log + string(trackJson) + "\n"
		id := int(track["id"].(float64))
		subtitleTrack := types.SubtitleTrack{
			UUID:      uuid.NewString(),
			TrackID:   id,
			TrackLang: trackLang,
			TrackName: trackName,
		}

		subtitleTracks = append(subtitleTracks, subtitleTrack)
	}

	if len(subtitleTracks) > 0 {
		extractCmd := os.Getenv("SUBTLER_EXTRACT")
		extractTemp := os.Getenv("SUBTLER_TEMP")
		args := []string{"tracks", video}
		for _, subtitleTrack := range subtitleTracks {
			id := subtitleTrack.TrackID
			uuid := subtitleTrack.UUID
			arg := fmt.Sprintf("%d:%s%c%s", id, extractTemp, os.PathSeparator, uuid)
			args = append(args, arg)
		}

		log = log + fmt.Sprintf("%s %s\n", extractCmd, args)
		run = exec.Command(extractCmd, args...)
		output, err = run.CombinedOutput()
		if err != nil {
			return log, err
		}

		log = log + fmt.Sprintf("%s\n", output)

		base := strings.TrimSuffix(video, ".mkv")
		for _, subtitleTrack := range subtitleTracks {
			uuid := subtitleTrack.UUID
			lang := subtitleTrack.TrackLang
			name := subtitleTrack.TrackName
			convertCmd := os.Getenv("SUBTLER_CONVERT")
			temp := fmt.Sprintf("%s%c%s", extractTemp, os.PathSeparator, uuid)
			srt := fmt.Sprintf("%s.%s.%s.srt", base, name, lang)
			args := []string{"-i", temp, srt}
			log = log + fmt.Sprintf("%s %s\n", convertCmd, args)
			run = exec.Command(convertCmd, args...)
			output, _ = run.CombinedOutput()
			log = log + fmt.Sprintf("%s\n", output)
		}
	}

	return log, nil
}
