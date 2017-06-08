# 手遊陰陽師(台服)onmyoji 事件通知器

透過 telegram gui界面，可以註冊手遊陰陽師(台服)的事件

## Install

`go get github.com/syhlion/onmyoji_event_bot`

## Usage

`ONMYOJI_EVENT_BOT_TOKEN={TOKEN} ./onmyoji_event_bot `

## Public Bot

* 已有在 telegram 註冊一組bot，搜尋 onmyoji_event_bot 或[點此](https://telegram.me/onmyoji_event_bot)，即可使用。

## Events

目前有通知的事件如下,可個別選擇事件註冊

1. 鬥技
2. 妖怪退治
3. 鬼王來襲
4. 陰界之門
5. 協同鬥技

## How did I create a bot ?

[Telegram Bot Document](https://core.telegram.org/bots#3-how-do-i-create-a-bot)

## Docker

[docker hub](https://hub.docker.com/syhlion/onmyoji_event_bot)

## Docker-Volumes:

* /onmyoji_bot/db

## Docker-ENV:

* ONMYOJI_EVENT_BOT_TOKEN

