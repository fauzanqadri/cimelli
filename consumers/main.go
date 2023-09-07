package main

import (
	"consumers/userinterestedoncreate"
	"consumers/userinterestedondelete"
	"consumers/userinterestedonupdate"
	"context"

	s "github.com/thejerf/suture/v4"
)

func main() {
	interestedoncreate := userinterestedoncreate.NewSuperVisor()
	interestedondelete := userinterestedondelete.NewSuperVisor()
	interestedonupdate := userinterestedonupdate.NewSuperVisor()

	spv := s.NewSimple("KafkaConsumers")

	spv.Add(interestedoncreate)
	spv.Add(interestedondelete)
	spv.Add(interestedonupdate)

	ctx := context.Background()

	spv.Serve(ctx)

}
