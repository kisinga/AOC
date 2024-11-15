// swift-tools-version: 6.0
public protocol Solvable {
    func solve() async throws -> Int
}

public protocol Logger {
    /// Log a message if verbose is enabled
    func log(_ message: String)
    /// Log a message regardless of verbose
    func forceLog(_ message: String)
}

// Basic digit extraction
extension String {
    // Extract all digits, preserving order
    func extractDigits() -> String {
        return self.filter { $0.isNumber }
    }

    func extractDigits() -> [Digit] {
        return self.enumerated().compactMap { index, element in
            if element.isNumber {
                return Digit(digit: element.wholeNumberValue!, position: index)
            }
            return nil
        }
    }
}

enum DigitMap: Int, CaseIterable {
    case zero = 0
    case one
    case two
    case three
    case four
    case five
    case six
    case seven
    case eight
    case nine

    var stringValue: String {
        String(describing: self)
    }
}

struct Digit {
    let value: Int
    let position: Int
    init(digit: Int, position: Int) {
        self.value = digit
        self.position = position
    }
}

extension String {
    func findPositions(of substring: String, options: String.CompareOptions = []) -> [Int] {
        var positions: [Int] = []
        var searchStartIndex = self.startIndex

        while let range = self.range(
            of: substring,
            options: options,
            range: searchStartIndex..<self.endIndex)
        {
            let position = distance(from: startIndex, to: range.lowerBound)
            positions.append(position)
            searchStartIndex = range.upperBound
        }

        return positions
    }

    func extractSpeltDigits() -> [Digit] {
        var occurrences: [Digit] = []
        for digit in DigitMap.allCases {
            let positions = self.findPositions(of: digit.stringValue, options: .caseInsensitive)
            if positions.capacity > 0 {
                occurrences.append(
                    contentsOf: positions.map { Digit(digit: digit.rawValue, position: $0) })
            }
        }
        return occurrences
    }
}
