package helpers

func GetDB[D any](datasource any) D {
	return datasource.(D)
}
