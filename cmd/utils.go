package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/exec"
	"strings"
)

func latestSubVersion(mainVersion string) string {
	result := make(map[string][]string)
	getVersionJsonFile(fmt.Sprintf(versionFileUrlTemplate, currentVersion), result)
	if len(result[mainVersion]) != 0 {
		return result[mainVersion][0]
	}

	log.Fatal("unknown dotnet core version, please use dvm listAll to check the version:", mainVersion, " exists or not")
	return ""
}

func getUsingVersion() string {
	path := getDotnetSdkPath("")
	cmd := exec.Command("ls", path)

	res, err := cmd.Output()

	if err != nil {
		log.Fatal("Can't get current using version, ", err, contactUs)
	}

	return string(res)
}

func getLatestDvmVersion() string {
	url := latestDvmVersionURL
	result := make(map[string][]string)
	getVersionJsonFile(url, result)
	return result["version"][0]
}

func removeOldDvmVersion() {
	oldDvmFile := fmt.Sprint(getDvmHome(), "/dvm")
	cmd := exec.Command("rm", oldDvmFile)
	if err := cmd.Run(); err != nil {
		log.Fatal("[Error] remove old dvm version error,", err, contactUs)
	}
}

func installLatestDvmVersion(version string) {
	resp, err := http.Get(fmt.Sprintf(DVM_COMMAND_REPO, version))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal("[Error] unexpected http GET status: ", resp.Status)
	}

	out, err := os.Create(fmt.Sprint(getDvmHome(), "/dvm"))

	if err != nil {
		log.Fatal("[Error] create dvm file error,", err, contactUs)
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)

	if err != nil {
		log.Fatal("[Error] copy file to dvm error", err, contactUs)
	}
	cmd := exec.Command("chmod", "+x", fmt.Sprint(getDvmHome(), "/dvm"))
	err = cmd.Run()
	if err != nil {
		log.Fatal("[Error] change dvm file permission error, now you can use `chmod +x ~/.dvm/dvm` to simply solve this", err, contactUs)
	}
}

func getVersionJsonFile(url string, result map[string][]string) {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatal("unexpected http GET status: ", resp.Status)
	}

	decodeError := json.NewDecoder(resp.Body).Decode(&result)
	if decodeError != nil {
		log.Fatal("cannot decode JSON: ", err)
	}
}

func checkSdkMainVersionExists(version string) bool {
	return checkPathExists(getDvmSdkStorePath(version)) || checkPathExists(getDotnetSdkPath(version))
}

func checkSdkSubVersionExists(version string) bool {
	subVersion := latestSubVersion(version)
	if len(subVersion) == 0 {
		os.Exit(1)
	}
	return checkPathExists(getDvmSdkStorePath(subVersion)) || checkPathExists(getDotnetSdkPath(subVersion))
}

func checkPathExists(path string) bool {
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true
	}
	return false
}

func getDvmHome() string {
	dvmHome := os.Getenv("DVM_HOME")
	if len(dvmHome) <= 0 {
		dvmHome = fmt.Sprint(os.Getenv("HOME"), "/.dvm")
	}
	return dvmHome
}

func getDotnetHome() string {

	dotnetHome := os.Getenv("DOTNET_HOME")
	if len(dotnetHome) <= 0 {
		dotnetHome = fmt.Sprint(os.Getenv("HOME"), "/.dotnet")
	}
	return dotnetHome
}

func removeOtherLink() {
	args := []string{"-rf", fmt.Sprint(getDotnetHome(), "/sdk/")}
	cmd := exec.Command("rm", args...)

	if err := cmd.Run(); err != nil {
		log.Fatal("remove other link failed when remove other link. ", err, contactUs)
	}
	createDir := exec.Command("mkdir", "-p", fmt.Sprint(getDotnetHome(), "/sdk"))

	if err := createDir.Run(); err != nil {
		log.Fatal("remove other link failed when create dir sdk. ", err, contactUs)
	}
}

func createLink(version string) {
	source := getDvmSdkStorePath(version)
	dest := getDotnetSdkPath("")
	cmd := exec.Command("ln", "-s", source, dest)
	err := cmd.Run()
	if err != nil {
		log.Fatal("create link failed when crate link. ", err, contactUs)
	}
}

func moveDotnetVersion(version string) {
	descPath := getDvmSdkStorePath("")
	sourcePath := getDotnetSdkPath(version)
	if moveFile(sourcePath, descPath) != nil {
		if !checkPathExists(getDvmSdkStorePath(version)) {
			log.Fatal("move dotnet version failed when move file check dvm home sdks. ", contactUs)
		}
	}
}

func getDotnetSdkPath(version string) string {
	if len(version) <= 0 {
		return fmt.Sprint(getDotnetHome(), "/sdk/")
	}
	return fmt.Sprint(getDotnetHome(), "/sdk/", version)
}

func getDvmSdkStorePath(version string) string {
	if len(version) <= 0 {
		return fmt.Sprint(getDvmHome(), "/sdks/")
	}
	return fmt.Sprint(getDvmHome(), "/sdks/", version)
}

func moveFile(sourceDir string, destDir string) error {
	args := []string{"-f", sourceDir, destDir}
	cmd := exec.Command("mv", args...)
	err := cmd.Run()
	return err
}

func getInstalledVersions() []string {
	sdksPath := getDvmSdkStorePath("")
	cmd := exec.Command("ls", sdksPath)
	out, err := cmd.Output()

	if err != nil {
		log.Fatal("Can't get installed versions, ", err, contactUs)
	}
	return strings.Split(string(out), "\n")
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
