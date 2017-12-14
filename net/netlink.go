package net

import (
	"github.com/pkg/errors"
	"github.com/vishvananda/netlink"
)

// Creates link if such does not exist. Prone to TOCTTOU.
func LinkAddIfNotExist(link netlink.Link) error {
	name := link.Attrs().Name
	if _, err := netlink.LinkByName(name); err != nil {
		if _, ok := err.(netlink.LinkNotFoundError); ok {
			if err = netlink.LinkAdd(link); err != nil {
				return errors.Wrapf(err, "netlink.LinkAdd %q", name)
			}
		} else {
			return errors.Wrapf(err, "netlink.LinkByName %q", name)
		}
	}

	return nil
}
