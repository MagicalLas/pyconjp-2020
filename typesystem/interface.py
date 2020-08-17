import abc


class Speaker:

    @abc.abstractmethod
    def speak(self) -> None:
        pass


class MySpeaker(Speaker):
    
    def speak(self) -> None:
        print("Las is talking")
