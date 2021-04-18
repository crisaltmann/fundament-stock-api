package service

type Service struct {

}

func (s Service) GetAllAssets() []Asset {
	return []Asset{
		Asset{1, "WEGE3", "WEG"},
		Asset{2, "ITUB3", "ITAÚ"},
	}
}

