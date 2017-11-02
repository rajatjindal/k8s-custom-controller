package discovery

type Discovery struct {
	Url string `json:"discoveryUrl"`
	Environment string `json:"environment"`
}

func (s *Discovery) ObjectCreated(obj interface{}) {
	updateDiscovery(s, obj, "created")
}

func (s *Discovery) ObjectDeleted(obj interface{}) {
	updateDiscovery(s, obj, "deleted")
}

func (s *Discovery) ObjectUpdated(oldObj, newObj interface{}) {
	updateDiscovery(s, newObj, "updated")
}

func updateDiscovery(s *Discovery, obj interface{}, action string) {

}

