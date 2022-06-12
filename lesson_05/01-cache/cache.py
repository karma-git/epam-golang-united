from typing import Any
from datetime import datetime


def new_object(key: str, value: Any, deadline: None | datetime = None) -> dict:
    return {key: {"value": value, "deadline": deadline}}


class Cache:
    def __init__(self) -> None:
        self._bucket = {}

    def _is_expired(self, obj: dict) -> bool:
        if obj.get("deadline") is None:
            return False
        elif obj.get("deadline") > datetime.now():
            return False
        else:
            return True

    def get(self, key: str) -> Any:
        result = self._bucket.get(key)
        if result and not self._is_expired(result):
            return result["value"]

    def put(self, obj: dict) -> None:
        self._bucket.update(obj)

    def keys(self) -> list[str]:
        a = [k for k in self._bucket.keys() if not self._is_expired(self._bucket[k])]
        return a

    def __repr__(self) -> list:
        return self._bucket.keys()

    def __str__(self) -> str:
        return ";".join(self._bucket.keys())


def main() -> None:
    c = Cache()
    obj1 = new_object("foo", "bar", datetime.fromisoformat("2021-05-05"))
    c.put(obj1)
    print(c.keys())
    print(c)


if __name__ == "__main__":
    main()
