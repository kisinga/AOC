package main

import (
	"fmt"
)

func main() {

	//This stores the array of expenses
	expenses := []int{1788, 1627, 1883, 1828, 1924, 1993, 972, 1840, 1866, 1762, 1781, 1782, 1520, 1971, 1660, 1857, 1867, 1564, 1983, 1391, 2002, 494, 1500, 1967, 1702, 1958, 1886, 1910, 1838, 1985, 1836, 2009, 2005, 1602, 1939, 1945, 1609, 1582, 1647, 1737, 1982, 1931, 790, 745, 1598, 1586, 1547, 1951, 1264, 1382, 1776, 1499, 1977, 1766, 1360, 1807, 1991, 1981, 1693, 634, 1847, 1774, 1990, 1409, 1410, 1974, 1862, 1744, 1827, 1978, 1980, 2003, 1491, 1595, 1640, 1576, 1887, 1746, 1617, 1923, 1706, 1964, 60, 1620, 1959, 257, 1395, 1854, 1843, 1682, 1667, 1639, 279, 1911, 1986, 1575, 1232, 1919, 1852, 1509, 1976, 1465, 2008, 1953, 1518, 1795, 1912, 1269, 1835, 1984, 1538, 2001, 1954, 1365, 1569, 1418, 1844, 1580, 1875, 1551, 1861, 1946, 1810, 1655, 1987, 1549, 1301, 1859, 1929, 1254, 1604, 1933, 1998, 1661, 1899, 1411, 1975, 1707, 1966, 1601, 1936, 1440, 1942, 1937, 1851, 1731, 1257, 1533, 1405, 1890, 1600, 1970, 1626, 1824, 1442, 2006, 1796, 1658, 1930, 646, 1904, 1489, 2004, 1922, 1424, 1802, 1623, 1870, 1242, 1591, 1338, 754, 1826, 1305, 1825, 1793, 1872, 1741, 1979, 107, 1833, 1856, 1952, 1791, 1728, 2010, 1965, 1646, 1522, 1671, 1624, 1876, 1537, 1759, 1962, 1773, 1907, 1573, 1908, 1903}
	//These are the correct values that add up to 2020
	var val1, val2, val3 int
	//used to break out of the loop, makes the code more efficient
	done := false
	//assign a var to the outer loop so that we can easily break when we find the right combination
outside:
	for i, expense := range expenses {
		//since the outer loop cannot find the values without relying on the inner loop, always check if the inner loop
		//had found the correct combination before proceeding to the next iteration, and conditionally break
		if done {
			break
		}
		for ii := i + 1; ii < len(expenses); ii++ {
			if expense+expenses[ii] > 2020 {
				continue
			}
			for _, val := range expenses {
				if expense+expenses[ii]+val == 2020 {
					val1, val2, val3 = expense, expenses[ii], val
					break outside
				}
			}
		}
	}
	//multiply the result as per the requirement
	fmt.Println(val1 * val2 * val3)
}
