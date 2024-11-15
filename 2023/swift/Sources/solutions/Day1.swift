// swift-tools-version: 6.0

import Foundation
import Types
import Utils

@frozen public enum Day1Errors: Error {
    case invalidDigitsLength
    case invalidPart
}

class Day1: Solvable {
    var filePath: String
    let logger: Logger
    let part: Int

    init(part: Int, filePath: String, logger: Logger) {
        self.filePath = filePath
        self.logger = logger
        self.part = part
    }

}

extension Day1 {
    func solve() async throws -> Int {
        self.logger.log("Solving day 1 Part \(part)")
        switch part {
        case 1:
            var sum = 0
            for try await line in FileReader.lines(from: filePath) {
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
                logger.log("Line sum: \(lineDigits)")
                sum += lineDigits
            }
            return sum

        case 2:
            var sum = 0
            for try await line in FileReader.lines(from: filePath) {
                let combinedDigits: [Digit] = line.extractDigits() + line.extractSpeltDigits()
                var firstDigit = Digit(digit: 0, position: line.count)
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
                logger.log("combinedDigits \(combinedDigits)")
                logger.log("Line sum: \(lineDigits)")
                sum += lineDigits
            }
            return sum
        default:
            preconditionFailure("Invalid part")
        }
    }
}
