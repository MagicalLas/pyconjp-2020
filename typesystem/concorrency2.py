import asyncio


async def hello():
    print("Hello Las")


async def main():
    await hello()
    print("Hello Pycon")


asyncio.run(main())
