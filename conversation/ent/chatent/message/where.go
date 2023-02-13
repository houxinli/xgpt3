// Code generated by ent, DO NOT EDIT.

package message

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/fanchunke/xgpt3/conversation/ent/chatent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldID, id))
}

// SessionID applies equality check predicate on the "session_id" field. It's identical to SessionIDEQ.
func SessionID(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSessionID, v))
}

// FromUserID applies equality check predicate on the "from_user_id" field. It's identical to FromUserIDEQ.
func FromUserID(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldFromUserID, v))
}

// ToUserID applies equality check predicate on the "to_user_id" field. It's identical to ToUserIDEQ.
func ToUserID(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldToUserID, v))
}

// Content applies equality check predicate on the "content" field. It's identical to ContentEQ.
func Content(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// SpouseID applies equality check predicate on the "spouse_id" field. It's identical to SpouseIDEQ.
func SpouseID(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSpouseID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// SessionIDEQ applies the EQ predicate on the "session_id" field.
func SessionIDEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSessionID, v))
}

// SessionIDNEQ applies the NEQ predicate on the "session_id" field.
func SessionIDNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSessionID, v))
}

// SessionIDIn applies the In predicate on the "session_id" field.
func SessionIDIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSessionID, vs...))
}

// SessionIDNotIn applies the NotIn predicate on the "session_id" field.
func SessionIDNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSessionID, vs...))
}

// SessionIDIsNil applies the IsNil predicate on the "session_id" field.
func SessionIDIsNil() predicate.Message {
	return predicate.Message(sql.FieldIsNull(FieldSessionID))
}

// SessionIDNotNil applies the NotNil predicate on the "session_id" field.
func SessionIDNotNil() predicate.Message {
	return predicate.Message(sql.FieldNotNull(FieldSessionID))
}

// FromUserIDEQ applies the EQ predicate on the "from_user_id" field.
func FromUserIDEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldFromUserID, v))
}

// FromUserIDNEQ applies the NEQ predicate on the "from_user_id" field.
func FromUserIDNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldFromUserID, v))
}

// FromUserIDIn applies the In predicate on the "from_user_id" field.
func FromUserIDIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldFromUserID, vs...))
}

// FromUserIDNotIn applies the NotIn predicate on the "from_user_id" field.
func FromUserIDNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldFromUserID, vs...))
}

// FromUserIDGT applies the GT predicate on the "from_user_id" field.
func FromUserIDGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldFromUserID, v))
}

// FromUserIDGTE applies the GTE predicate on the "from_user_id" field.
func FromUserIDGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldFromUserID, v))
}

// FromUserIDLT applies the LT predicate on the "from_user_id" field.
func FromUserIDLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldFromUserID, v))
}

// FromUserIDLTE applies the LTE predicate on the "from_user_id" field.
func FromUserIDLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldFromUserID, v))
}

// FromUserIDContains applies the Contains predicate on the "from_user_id" field.
func FromUserIDContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldFromUserID, v))
}

// FromUserIDHasPrefix applies the HasPrefix predicate on the "from_user_id" field.
func FromUserIDHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldFromUserID, v))
}

// FromUserIDHasSuffix applies the HasSuffix predicate on the "from_user_id" field.
func FromUserIDHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldFromUserID, v))
}

// FromUserIDEqualFold applies the EqualFold predicate on the "from_user_id" field.
func FromUserIDEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldFromUserID, v))
}

// FromUserIDContainsFold applies the ContainsFold predicate on the "from_user_id" field.
func FromUserIDContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldFromUserID, v))
}

// ToUserIDEQ applies the EQ predicate on the "to_user_id" field.
func ToUserIDEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldToUserID, v))
}

// ToUserIDNEQ applies the NEQ predicate on the "to_user_id" field.
func ToUserIDNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldToUserID, v))
}

// ToUserIDIn applies the In predicate on the "to_user_id" field.
func ToUserIDIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldToUserID, vs...))
}

// ToUserIDNotIn applies the NotIn predicate on the "to_user_id" field.
func ToUserIDNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldToUserID, vs...))
}

// ToUserIDGT applies the GT predicate on the "to_user_id" field.
func ToUserIDGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldToUserID, v))
}

// ToUserIDGTE applies the GTE predicate on the "to_user_id" field.
func ToUserIDGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldToUserID, v))
}

// ToUserIDLT applies the LT predicate on the "to_user_id" field.
func ToUserIDLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldToUserID, v))
}

// ToUserIDLTE applies the LTE predicate on the "to_user_id" field.
func ToUserIDLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldToUserID, v))
}

// ToUserIDContains applies the Contains predicate on the "to_user_id" field.
func ToUserIDContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldToUserID, v))
}

// ToUserIDHasPrefix applies the HasPrefix predicate on the "to_user_id" field.
func ToUserIDHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldToUserID, v))
}

// ToUserIDHasSuffix applies the HasSuffix predicate on the "to_user_id" field.
func ToUserIDHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldToUserID, v))
}

// ToUserIDEqualFold applies the EqualFold predicate on the "to_user_id" field.
func ToUserIDEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldToUserID, v))
}

// ToUserIDContainsFold applies the ContainsFold predicate on the "to_user_id" field.
func ToUserIDContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldToUserID, v))
}

// ContentEQ applies the EQ predicate on the "content" field.
func ContentEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldContent, v))
}

// ContentNEQ applies the NEQ predicate on the "content" field.
func ContentNEQ(v string) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldContent, v))
}

// ContentIn applies the In predicate on the "content" field.
func ContentIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldContent, vs...))
}

// ContentNotIn applies the NotIn predicate on the "content" field.
func ContentNotIn(vs ...string) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldContent, vs...))
}

// ContentGT applies the GT predicate on the "content" field.
func ContentGT(v string) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldContent, v))
}

// ContentGTE applies the GTE predicate on the "content" field.
func ContentGTE(v string) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldContent, v))
}

// ContentLT applies the LT predicate on the "content" field.
func ContentLT(v string) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldContent, v))
}

// ContentLTE applies the LTE predicate on the "content" field.
func ContentLTE(v string) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldContent, v))
}

// ContentContains applies the Contains predicate on the "content" field.
func ContentContains(v string) predicate.Message {
	return predicate.Message(sql.FieldContains(FieldContent, v))
}

// ContentHasPrefix applies the HasPrefix predicate on the "content" field.
func ContentHasPrefix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasPrefix(FieldContent, v))
}

// ContentHasSuffix applies the HasSuffix predicate on the "content" field.
func ContentHasSuffix(v string) predicate.Message {
	return predicate.Message(sql.FieldHasSuffix(FieldContent, v))
}

// ContentEqualFold applies the EqualFold predicate on the "content" field.
func ContentEqualFold(v string) predicate.Message {
	return predicate.Message(sql.FieldEqualFold(FieldContent, v))
}

// ContentContainsFold applies the ContainsFold predicate on the "content" field.
func ContentContainsFold(v string) predicate.Message {
	return predicate.Message(sql.FieldContainsFold(FieldContent, v))
}

// SpouseIDEQ applies the EQ predicate on the "spouse_id" field.
func SpouseIDEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldSpouseID, v))
}

// SpouseIDNEQ applies the NEQ predicate on the "spouse_id" field.
func SpouseIDNEQ(v int) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldSpouseID, v))
}

// SpouseIDIn applies the In predicate on the "spouse_id" field.
func SpouseIDIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldSpouseID, vs...))
}

// SpouseIDNotIn applies the NotIn predicate on the "spouse_id" field.
func SpouseIDNotIn(vs ...int) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldSpouseID, vs...))
}

// SpouseIDIsNil applies the IsNil predicate on the "spouse_id" field.
func SpouseIDIsNil() predicate.Message {
	return predicate.Message(sql.FieldIsNull(FieldSpouseID))
}

// SpouseIDNotNil applies the NotNil predicate on the "spouse_id" field.
func SpouseIDNotNil() predicate.Message {
	return predicate.Message(sql.FieldNotNull(FieldSpouseID))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Message {
	return predicate.Message(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Message {
	return predicate.Message(sql.FieldLTE(FieldCreatedAt, v))
}

// HasSpouse applies the HasEdge predicate on the "spouse" edge.
func HasSpouse() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSpouseWith applies the HasEdge predicate on the "spouse" edge with a given conditions (other predicates).
func HasSpouseWith(preds ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2O, false, SpouseTable, SpouseColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasSession applies the HasEdge predicate on the "session" edge.
func HasSession() predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SessionTable, SessionColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasSessionWith applies the HasEdge predicate on the "session" edge with a given conditions (other predicates).
func HasSessionWith(preds ...predicate.Session) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(SessionInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, SessionTable, SessionColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Message) predicate.Message {
	return predicate.Message(func(s *sql.Selector) {
		p(s.Not())
	})
}
