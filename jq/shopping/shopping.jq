.name,
(.ingredients | length),
(.ingredients[] | select(.item=="sugar") | .amount.quantity),
(.ingredients + ."optional ingredients" | map(select(.substitute) | {(.item): .substitute}) | add)
