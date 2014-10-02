package sensu

import "fmt"

// GetClients Return the list of clients
func (s *Sensu) GetClients(limit int, offset int) ([]interface{}, error) {
	return s.GetList("clients", limit, offset)
}

// GetClient Return client info
func (s *Sensu) GetClient(client string) ([]interface{}, error) {
	return s.Get(fmt.Sprintf("client/%s", client))
}

// GetClientHistory Return client history
func (s *Sensu) GetClientHistory(client string) ([]interface{}, error) {
	return s.Get(fmt.Sprintf("client/%s/history", client))
}

// DeleteClient Return the list of clients
func (s *Sensu) DeleteClient(client string) ([]interface{}, error) {
	return s.Delete(fmt.Sprintf("client/%s", client))
}