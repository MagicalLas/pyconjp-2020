class Speaker(object):
    def __init__(self, name: str, age: int):
        self.name = name
        self._age = age

    def speak(self) -> None:
        print(f"now {self.name} is speaking")

las = Speaker("las", 20)
las.speak()
