package entities

func ResourceOfferFromRU(ru ResourceUnits) ResourcesOffer {
	res := make(ResourcesOffer, 0, len(ru))

	for _, r := range ru {
		res = append(res, ResourceOffer{
			Resources: r.Resources,
			Count:     r.Count,
		})
	}

	return res
}
