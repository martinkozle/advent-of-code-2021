@ECHO OFF
set SESSION_ID=%1
set _tail=%*
call set _tail=%%_tail:*%1=%%
echo %_tail%
python -m aoc_to_markdown %_tail%
