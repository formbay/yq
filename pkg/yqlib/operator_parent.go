package yqlib

import "container/list"

func getParentOperator(_ *dataTreeNavigator, context Context, _ *ExpressionNode) (Context, error) {
	log.Debugf("-- getParentOperator")

	var results = list.New()

	for el := context.MatchingNodes.Front(); el != nil; el = el.Next() {
		candidate := el.Value.(*CandidateNode)
		if candidate.Parent != nil {
			results.PushBack(candidate.Parent)
		}
	}

	return context.ChildContext(results), nil

}
