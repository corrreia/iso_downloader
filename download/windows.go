package download

func (d Download) GetWindows11() ([]string, error) {
	return []string{"https://software.download.prss.microsoft.com/dbazure/Win11_22H2_English_x64v1.iso?t=6f756a97-c24e-4208-adee-4f99b83a39d3&e=1671018899&h=e2b805afde87e7b79eca27323bad265ba74449f00831ea8b5a31d4a62fb85fe3",
			"Windows",
			"11",
			"Windows"},
		nil
}

func (d Download) GetWindows10() ([]string, error) {
	return []string{"https://software.download.prss.microsoft.com/dbazure/Win10_22H2_English_x64.iso?t=dd8444dd-680c-42b7-94e6-d51522592ed9&e=1671018722&h=7312e04a37ab9f02973a3502be537db00d7ea1102df8683fb0e347ec9121f109",
			"Windows",
			"10",
			"Windows"},
		nil
}