<script lang="ts">
    import {_} from 'svelte-i18n'
    import PopupContainer from "../../PopupContainer/PopupContainer.svelte";
    import {style} from "./style";
    import {onMount} from "svelte";
    import {
        setCurrentSession,
        updateUserDataInStore,
        UserDataExtended
    } from "../../../../stores/auth";
    import {EntityType, ActionType} from "../../../../stores/utils";
    import UserSettingsForm from "../UserSettingsForm/UserSettingsForm.svelte";

    const onSaveHandler = () => {
        const error = updateUserDataInStore(data, initialData.login);

        if (error) {
            //backend validation
            return;
        }

        //TODO: update session by websocket from backend
        setCurrentSession(Object.assign({}, data, {
            password: undefined,
            newPassword: undefined
        }));
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: UserDataExtended = {};
    export let data: UserDataExtended = {};
    export let isPopupOpened: boolean = false;
</script>

<PopupContainer let:outerPopupUuid={uuid}
                popupClass={style.Popup} onAcceptHandler={onSaveHandler}
                onCancelHandler={onCloseHandler} withAcceptButton={true}
                withCancelButton={$_("common.components.buttons.cancel")}
                entityType={EntityType.UserSettings}
                actionType={ActionType.Edit} entityId={data.login}
                title={`${$_("user.titles.edit")}`}
                isPopupOpened={isPopupOpened}
                onCloseHandler={onCloseHandler}
                acceptButtonTitle={$_("common.components.buttons.save")}>
    <UserSettingsForm outerPopupUuid={uuid} data={data}/>
</PopupContainer>
