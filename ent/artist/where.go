// Code generated by entc, DO NOT EDIT.

package artist

import (
	"enttest/ent/predicate"

	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Artist {
	return predicate.Artist(
		func(s *sql.Selector) {
			s.Where(sql.EQ(s.C(FieldID), id))
		},
	)
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	},
	)
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	},
	)
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	},
	)
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	},
	)
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	},
	)
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	},
	)
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	},
	)
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	},
	)
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	},
	)
}

// Age applies equality check predicate on the "age" field. It's identical to AgeEQ.
func Age(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	},
	)
}

// Phone applies equality check predicate on the "phone" field. It's identical to PhoneEQ.
func Phone(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	},
	)
}

// Plays applies equality check predicate on the "plays" field. It's identical to PlaysEQ.
func Plays(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlays), v))
	},
	)
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	},
	)
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	},
	)
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlug), v...))
	},
	)
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	},
	)
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	},
	)
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	},
	)
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	},
	)
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	},
	)
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	},
	)
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	},
	)
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	},
	)
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	},
	)
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	},
	)
}

// AgeEQ applies the EQ predicate on the "age" field.
func AgeEQ(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAge), v))
	},
	)
}

// AgeNEQ applies the NEQ predicate on the "age" field.
func AgeNEQ(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAge), v))
	},
	)
}

// AgeIn applies the In predicate on the "age" field.
func AgeIn(vs ...int) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAge), v...))
	},
	)
}

// AgeNotIn applies the NotIn predicate on the "age" field.
func AgeNotIn(vs ...int) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAge), v...))
	},
	)
}

// AgeGT applies the GT predicate on the "age" field.
func AgeGT(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAge), v))
	},
	)
}

// AgeGTE applies the GTE predicate on the "age" field.
func AgeGTE(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAge), v))
	},
	)
}

// AgeLT applies the LT predicate on the "age" field.
func AgeLT(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAge), v))
	},
	)
}

// AgeLTE applies the LTE predicate on the "age" field.
func AgeLTE(v int) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAge), v))
	},
	)
}

// PhoneEQ applies the EQ predicate on the "phone" field.
func PhoneEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPhone), v))
	},
	)
}

// PhoneNEQ applies the NEQ predicate on the "phone" field.
func PhoneNEQ(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPhone), v))
	},
	)
}

// PhoneIn applies the In predicate on the "phone" field.
func PhoneIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPhone), v...))
	},
	)
}

// PhoneNotIn applies the NotIn predicate on the "phone" field.
func PhoneNotIn(vs ...string) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPhone), v...))
	},
	)
}

// PhoneGT applies the GT predicate on the "phone" field.
func PhoneGT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPhone), v))
	},
	)
}

// PhoneGTE applies the GTE predicate on the "phone" field.
func PhoneGTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPhone), v))
	},
	)
}

// PhoneLT applies the LT predicate on the "phone" field.
func PhoneLT(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPhone), v))
	},
	)
}

// PhoneLTE applies the LTE predicate on the "phone" field.
func PhoneLTE(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPhone), v))
	},
	)
}

// PhoneContains applies the Contains predicate on the "phone" field.
func PhoneContains(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPhone), v))
	},
	)
}

// PhoneHasPrefix applies the HasPrefix predicate on the "phone" field.
func PhoneHasPrefix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPhone), v))
	},
	)
}

// PhoneHasSuffix applies the HasSuffix predicate on the "phone" field.
func PhoneHasSuffix(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPhone), v))
	},
	)
}

// PhoneIsNil applies the IsNil predicate on the "phone" field.
func PhoneIsNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldPhone)))
	},
	)
}

// PhoneNotNil applies the NotNil predicate on the "phone" field.
func PhoneNotNil() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldPhone)))
	},
	)
}

// PhoneEqualFold applies the EqualFold predicate on the "phone" field.
func PhoneEqualFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPhone), v))
	},
	)
}

// PhoneContainsFold applies the ContainsFold predicate on the "phone" field.
func PhoneContainsFold(v string) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPhone), v))
	},
	)
}

// PlaysEQ applies the EQ predicate on the "plays" field.
func PlaysEQ(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPlays), v))
	},
	)
}

// PlaysNEQ applies the NEQ predicate on the "plays" field.
func PlaysNEQ(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPlays), v))
	},
	)
}

// PlaysIn applies the In predicate on the "plays" field.
func PlaysIn(vs ...int64) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPlays), v...))
	},
	)
}

// PlaysNotIn applies the NotIn predicate on the "plays" field.
func PlaysNotIn(vs ...int64) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPlays), v...))
	},
	)
}

// PlaysGT applies the GT predicate on the "plays" field.
func PlaysGT(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPlays), v))
	},
	)
}

// PlaysGTE applies the GTE predicate on the "plays" field.
func PlaysGTE(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPlays), v))
	},
	)
}

// PlaysLT applies the LT predicate on the "plays" field.
func PlaysLT(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPlays), v))
	},
	)
}

// PlaysLTE applies the LTE predicate on the "plays" field.
func PlaysLTE(v int64) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPlays), v))
	},
	)
}

// GenderEQ applies the EQ predicate on the "gender" field.
func GenderEQ(v Gender) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldGender), v))
	},
	)
}

// GenderNEQ applies the NEQ predicate on the "gender" field.
func GenderNEQ(v Gender) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldGender), v))
	},
	)
}

// GenderIn applies the In predicate on the "gender" field.
func GenderIn(vs ...Gender) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldGender), v...))
	},
	)
}

// GenderNotIn applies the NotIn predicate on the "gender" field.
func GenderNotIn(vs ...Gender) predicate.Artist {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Artist(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(vs) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldGender), v...))
	},
	)
}

// HasSongs applies the HasEdge predicate on the "songs" edge.
func HasSongs() predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SongsTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SongsTable, SongsColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	},
	)
}

// HasSongsWith applies the HasEdge predicate on the "songs" edge with a given conditions (other predicates).
func HasSongsWith(preds ...predicate.Song) predicate.Artist {
	return predicate.Artist(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SongsInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, SongsTable, SongsColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	},
	)
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for _, p := range predicates {
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Artist) predicate.Artist {
	return predicate.Artist(
		func(s *sql.Selector) {
			s1 := s.Clone().SetP(nil)
			for i, p := range predicates {
				if i > 0 {
					s1.Or()
				}
				p(s1)
			}
			s.Where(s1.P())
		},
	)
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Artist) predicate.Artist {
	return predicate.Artist(
		func(s *sql.Selector) {
			p(s.Not())
		},
	)
}