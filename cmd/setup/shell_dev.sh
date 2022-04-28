# #DEV_START#
# #DESCRIPTION#
dev() {
    export #ENV_VAR#=#HOME#/~/.lazygit/newdir

    lazygit "$@"

    if [ -f $#ENV_VAR# ]; then
        cd "$(cat $#ENV_VAR#)"
        rm -f $#ENV_VAR# >/dev/null
    fi
}
# #DEV_END#