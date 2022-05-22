from typing import Tuple


class errorEmptyInput(Exception):
    pass


class errorNotTwoOperands(Exception):
    pass


def string_sum(
    user_input: str,
) -> Tuple[str, None | errorEmptyInput | errorNotTwoOperands]:
    digit_count = lambda user_input: len([d for d in user_input if d.isdecimal()])
    if user_input.strip() == "" or digit_count(user_input) == 0:
        # raise errorEmptyInput
        return "", errorEmptyInput
    elif digit_count(user_input) != 2:
        # raise errorNotTwoOperands
        return "", errorNotTwoOperands
    else:
        try:
            # replace all whitespace characters from string
            user_input = "".join(user_input.split())
            result = eval(user_input)
        except Exception as e:
            return "", e
        else:
            return result, None


def main() -> None:
    user_input = input()
    result, err = string_sum(user_input)
    print(result, err)


if __name__ == "__main__":
    main()
