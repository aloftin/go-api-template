package models

import "fmt"

// ResourceNotFoundError represents an error that is thrown because a resource could not be found
type ResourceNotFoundError struct {
	ResourceName string
}

func (nfe *ResourceNotFoundError) Error() string {
	if nfe.ResourceName == "" {
		return "Resource not found"
	}

	return fmt.Sprintf("%s not found", nfe.ResourceName)
}

// InvalidInputError represents an error that is thrown because an operation did not receive valid input parameters
type InvalidInputError struct {
	Message string
}

func (b *InvalidInputError) Error() string {
	if b.Message == "" {
		return "Bad request"
	}

	return b.Message
}
