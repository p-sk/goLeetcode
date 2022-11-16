/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */
package main

func isBadVersion(n int) bool {

	return false
}

func firstBadVersion(n int) int {
	for n > 0 {
		if !isBadVersion(n) {
			n += 1
			return n
		} else {
			n -= 1
		}
	}
	return n + 1
}
