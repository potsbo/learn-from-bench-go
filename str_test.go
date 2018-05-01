package main

import (
	"testing"
)

var str = `
Lorem ipsum dolor sit amet, consectetur adipiscing elit. Vestibulum eu odio non eros porta elementum. Aliquam erat volutpat. Mauris a lectus malesuada, tincidunt purus a, consequat mauris. Aenean at ante odio. Sed gravida accumsan lacus, a pretium neque. Integer sit amet volutpat ipsum. Aliquam tincidunt justo at suscipit pharetra. Nunc ac hendrerit nisi. Sed venenatis massa at dignissim dictum. Aenean quis metus laoreet, consequat urna nec, pharetra ipsum. Suspendisse mollis eros id pulvinar tristique. Maecenas id dolor non erat pretium tincidunt.

Sed id neque a mauris ultricies aliquet. Nunc ullamcorper nunc eu orci imperdiet, vel pulvinar nisl commodo. Ut in augue ac lacus laoreet pretium. Nunc vel varius ligula, finibus laoreet magna. Fusce gravida elit ipsum, sit amet posuere justo viverra eu. Etiam lorem arcu, varius eu arcu eget, imperdiet volutpat ex. Duis ante ante, blandit id malesuada eget, pellentesque ut metus. Suspendisse pretium rutrum sapien sit amet dignissim. Morbi ullamcorper volutpat nunc, non laoreet est maximus vel. Pellentesque suscipit dignissim nibh at blandit. Nulla vel finibus velit. Fusce sit amet scelerisque dui. Fusce venenatis enim non nisi porttitor, interdum aliquet urna laoreet. Integer varius bibendum nulla, at porttitor odio sagittis at.

Integer dictum quam turpis. Cras ut purus at turpis vehicula elementum non in augue. Maecenas vel purus consequat, elementum justo ut, convallis ligula. Quisque tincidunt ut quam sed varius. Pellentesque pellentesque condimentum purus eu tempus. Vivamus pulvinar nulla augue, non cursus risus aliquam sit amet. Curabitur eget turpis ac diam tempus cursus in in risus. Donec dapibus lorem leo.

Cras vitae feugiat dui. Donec facilisis venenatis mollis. Pellentesque finibus, purus nec tempus tincidunt, sem risus rhoncus erat, nec euismod diam eros eget felis. Nam non mi at ex imperdiet tincidunt non et urna. Donec et leo sed mi luctus porttitor quis volutpat ante. Donec elementum, metus eu posuere molestie, arcu libero viverra urna, vel fringilla eros orci ut mauris. Donec volutpat commodo lorem quis rutrum. Suspendisse finibus sollicitudin metus sed dapibus. Duis vel commodo quam. Mauris vitae fringilla ipsum.

Donec diam quam, luctus non est non, euismod dignissim purus. In hac habitasse platea dictumst. Nulla ultrices consectetur sapien. Maecenas nulla dui, finibus vel turpis quis, hendrerit rhoncus dui. Morbi magna ex, sagittis sollicitudin interdum a, tincidunt sit amet dolor. Duis id egestas est. Curabitur a magna vehicula, blandit est a, ultrices lacus. Curabitur vitae purus rutrum, facilisis mauris in, dapibus velit. Ut non tristique nisi. Praesent sit amet velit neque. Quisque mattis orci vel vulputate lobortis. Ut eu sapien turpis. Phasellus porta elementum libero, ut mollis sem vestibulum nec. Donec faucibus egestas tincidunt. Donec accumsan fermentum sapien. Donec faucibus velit vitae tellus pulvinar, quis posuere eros iaculis.
`

func BenchmarkLen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = len(str) != 0
	}
}

func BenchmarkEmpty(b *testing.B) {
	for n := 0; n < b.N; n++ {
		_ = str != ""
	}
}
