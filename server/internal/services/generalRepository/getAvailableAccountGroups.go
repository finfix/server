package generalRepository

func (repo *GeneralRepository) GetAvailableAccountGroups(userID uint32) []uint32 {
	availableAccountGroupIDs := make([]uint32, 0, len(repo.accesses.Get()[userID]))
	for accountGroupID := range repo.accesses.Get()[userID] {
		availableAccountGroupIDs = append(availableAccountGroupIDs, accountGroupID)
	}
	return availableAccountGroupIDs
}
