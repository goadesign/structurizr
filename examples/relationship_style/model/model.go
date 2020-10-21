package design

import (
	. "goa.design/model/dsl"
)

var _ = Design("Getting Started", "This is a model of my software system.", func() {
	SoftwareSystem("The System", "My Software System")
	Person("Person1", "A person using the system.", func() {
		Uses("The System", "Edge\nwith vertices", func() {
			Tag("pos75")
		})
		Tag("person")
	})
	Person("Person2", "The intellectual.\nTwo relationships\nautomatically spread", func() {
		Uses("The System", "Left")
		Uses("The System", "Right")
		Tag("person")
	})
	Person("Person3", "Another person", func() {
		Uses("The System", "Solid Orthogonal", func() {
			Tag("knows")
		})
	})

	SoftwareSystem("System 2", "Software System 2")
	Person("Person4", "", func() {
		Tag("person")
		Uses("System 2", "O")
		Uses("The System", "O")
	})

	Views(func() {
		SystemContextView("The System", "SystemContext", "System Context diagram.", func() {
			AddAll()
			Link("Person1", "The System", "Edge\nwith vertices", func() {
				Vertices(300, 300, 300, 800)
			})
			AutoLayout(RankLeftRight)
		})
		Styles(func() {
			ElementStyle("person", func() {
				Shape(ShapePerson)
			})
			RelationshipStyle("pos75", func() {
				Position(40)
			})
			RelationshipStyle("knows", func() {
				Routing(RoutingOrthogonal)
				//Routing(RoutingDirect)
				Solid()
			})
		})
	})
})
