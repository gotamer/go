package mail

// How many keys do we have in the header
func (h Header) Len() int {
	return len(h)
}

// Has given key in the Header
func (h Header) Has(key string) (has bool) {
	_, has = h[key]
	return
}

// What keys do we have in the Header, preserving order
func (h Header) Keys() (keys []string) {
	for _, v := range h {
		keys = append(keys, v)
	}
	return
}

// Sometimes we need to delete a key such as "Bcc"
func (h Header) Del(key string) {
	delete(h, key)
}

// Append key to Header
func (h Header) Add(key string, value string) {
	values, _ := h[key]
	values = append(values, value)
	h[key] = values
}

// Received: from mail.wndp.co
//
//	by m.wndp.co with LMTP
//	id oDDXORaPqmVb5hIAPrcMpw
//	(envelope-from <patrick@example.com>)
//	for <hi@exaple.pw>; Fri, 19 Jan 2024 16:02:46 +0100
//
// Prepend to Received slice
func (h Header) PrependReceived(value string) {
	if hdr, has := h["Received"]; has {
		hdr = append([]string{value}, hdr...)
		h["Received"] = hdr
	} else {
		var newH = make(Header, len(h)) // map[string][]string
		newH["Received"] = []string{value}
		for k, v := range h {
			newH[k] = v
		}
		h = newH
	}
}
