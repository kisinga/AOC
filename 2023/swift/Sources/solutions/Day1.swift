// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum Day1Errors: Error {
    case invalidDigitsLength
    case invalidPart
}

struct Day1: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int
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

struct CustomString {
    let value: String
    init(value: String) {
        self.value = value
    }
}

extension CustomString {
    // Extract all digits, preserving order
    func extractDigits() -> String {
        return self.value.filter { $0.isNumber }
    }

    func extractDigits() -> [Digit] {
        return self.value.enumerated().compactMap { index, element in
            if element.isNumber {
                return Digit(digit: element.wholeNumberValue!, position: index)
            }
            return nil
        }
    }

    func findPositions(of substring: String, options: String.CompareOptions = []) -> [Int] {
        var positions: [Int] = []
        var searchStartIndex = self.value.startIndex

        while let range = self.value.range(
            of: substring,
            options: options,
            range: searchStartIndex..<self.value.endIndex)
        {
            let position = self.value.distance(from: self.value.startIndex, to: range.lowerBound)
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

extension Day1 {
    func solve() async throws -> Int {
        switch part {
        case 1:
            var sum = 0
            for try await lineString in FileReader.lines(from: filePath) {
                let line = CustomString(value: lineString)
                let digits: String = line.extractDigits()
                let lineDigits: Int

                switch digits.count {
                case 0:
                    lineDigits = 0
                case 1:
                    let digit = digits.first!.wholeNumberValue!
                    lineDigits = digit * 10 + digit
                case let x where x > 1:
                    let firstChar = digits.first
                    let lastChar = digits.last
                    if let firstChar, let lastChar {
                        lineDigits = firstChar.wholeNumberValue! * 10 + lastChar.wholeNumberValue!
                    } else {
                        throw Day1Errors.invalidDigitsLength
                    }
                default:
                    preconditionFailure("Invalid digits length. It should never happen")
                }
                logger.log(.debug("Line sum: \(lineDigits)"))
                sum += lineDigits
            }
            return sum

        case 2:
            var sum = 0
            for try await lineString in FileReader.lines(from: filePath) {
                let line = CustomString(value: lineString)
                let combinedDigits: [Digit] = line.extractDigits() + line.extractSpeltDigits()
                var firstDigit = Digit(digit: 0, position: line.value.count)
                var lastDigit = Digit(digit: 0, position: 0)
                combinedDigits.forEach { digit in
                    if digit.position <= firstDigit.position {
                        firstDigit = digit
                    }
                    if digit.position >= lastDigit.position {
                        lastDigit = digit
                    }
                }
                let lineDigits = (firstDigit.value * 10 + lastDigit.value)
                logger.log(.debug("combinedDigits \(combinedDigits)"))
                logger.log(.debug("Line sum: \(lineDigits)"))
                sum += lineDigits
            }
            return sum
        default:
            preconditionFailure("Invalid part")
        }
    }
}
