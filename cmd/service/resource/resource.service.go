package resourceservice

import (
	ResourceInterface "github.com/MohitVachhani/go-learn/pkg/structs/resource"

	ResourceRepository "github.com/MohitVachhani/go-learn/cmd/repo/resource"
)

func GetResourceById(input ResourceInterface.GetResourceInputInterface) ResourceInterface.ResourceSchema {

	resourceId := input.ResourceId

	resource := ResourceRepository.GetResourceById(resourceId)

	return resource
}
