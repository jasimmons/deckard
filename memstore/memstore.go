package memstore

import (
	"context"
	"errors"

	"github.com/jasimmons/deckard"
)

var (
	ErrIdentifierNotFound = errors.New("identifier not found")
	ErrMaxIdentifiers     = errors.New("maximum identifier limit reached")
	ErrIdentifierExists   = errors.New("identifier already exists")
)

type MemStore struct {
	maxIdentifiers int
	identifiers    map[int]*deckard.Identifier
}

func New(maxIdentifiers int) *MemStore {
	return &MemStore{
		maxIdentifiers: maxIdentifiers,
		identifiers:    make(map[int]*deckard.Identifier),
	}
}

func (s *MemStore) ListIdentifiers(ctx context.Context, tags ...string) ([]*deckard.Identifier, error) {
	filteredIdentifiers := make([]*deckard.Identifier, 0)
	for _, identifier := range s.identifiers {
		if identifier.HasTags(tags...) {
			filteredIdentifiers = append(filteredIdentifiers, identifier)
		}
	}

	return filteredIdentifiers, nil
}

func (s *MemStore) GetIdentifier(ctx context.Context, identifierId int) (*deckard.Identifier, error) {
	if _, ok := s.identifiers[identifierId]; !ok {
		return nil, ErrIdentifierNotFound
	}

	return s.identifiers[identifierId], nil
}

func (s *MemStore) CreateIdentifier(ctx context.Context, identifier *deckard.Identifier) (*deckard.Identifier, error) {
	if len(s.identifiers) >= s.maxIdentifiers {
		return nil, ErrMaxIdentifiers
	}

	if identifier.ID != 0 {
		if _, ok := s.identifiers[identifier.ID]; ok {
			return nil, ErrIdentifierExists
		}
	}

	// identifier ID is 0/empty, so we need to set an it
	for ctr := 1; ctr < s.maxIdentifiers; ctr++ {
		if _, ok := s.identifiers[ctr]; !ok {
			// found next open ID where no identifier exists
			identifier.ID = ctr
			s.identifiers[ctr] = identifier
			return identifier, nil
		}
	}

	return nil, errors.New("unknown error creating identifier")
}
