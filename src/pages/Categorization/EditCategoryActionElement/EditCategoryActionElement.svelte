<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {Category} from "../types";
    import Button
        from "../../../common/components/Buttons/Button/Button.svelte";
    import ButtonIconEdit
        from "../../../common/components/Buttons/ButtonIconEdit/ButtonIconEdit.svelte";
    import PopupContainer
        from "../../../common/components/PopupContainer/PopupContainer.svelte";
    import {Popup} from "./style";
    import {EntityType, ActionType} from "../../../stores/utils";
    import CategoryEditForm from "../CategoryEditForm/CategoryEditForm.svelte";
    import {
        addCategoryToStore,
        updateCategoryInStore
    } from "../../../stores/categories";

    let isPopupOpened = false;

    const onClickHandler = () => {
        isPopupOpened = true;
    }

    const onCloseHandler = () => {
        isPopupOpened = false;
        data = {...initialData}
    }

    const onSaveHandler = () => {
        console.log("onSaveHandler", data)
        if (data.id) {
            updateCategoryInStore(data)
        } else {
            addCategoryToStore(data)
        }
    }

    onMount(() => {
        initialData = {...data};
    })

    let initialData: Category = {};

    export let data: Category = {};
    export let withButton: boolean = false;
    export let buttonTitle: string = "";
</script>


<div>
    {#if withButton}
        <Button title={buttonTitle} onClickHandler={onClickHandler}/>
    {:else}
        <ButtonIconEdit width={16} height={16} onClickHandler={onClickHandler}/>
    {/if}
    {#if isPopupOpened}
        <PopupContainer let:outerPopupUuid={uuid}
                        popupClass={Popup} onAcceptHandler={onSaveHandler}
                        onCancelHandler={onCloseHandler} withAcceptButton={true}
                        withCancelButton={$_("common.components.buttons.cancel")}
                        entityType={EntityType.Category}
                        actionType={ActionType.Remove} entityId={data.id}
                        title={data.id ? `${$_("common.titles.edit")} ${data.name}` || $_("common.titles.edit") : $_("categories.titles.new")}
                        isPopupOpened={isPopupOpened}
                        onCloseHandler={onCloseHandler}
                        acceptButtonTitle={$_("common.components.buttons.save")}>
            <CategoryEditForm outerPopupUuid={uuid} data={data}/>
        </PopupContainer>
    {/if}

</div>