<script lang="ts">
    import {_} from 'svelte-i18n'
    import {onMount} from "svelte";
    import {Purchase} from "../types";
    import {getSimpleShopsData} from "../../../stores/shops";
    import {
        SideMinorPadding,
        TextArea,
        FlexVert,
        FlexHor,
        Form,
        MainFieldsWrapper
    } from "./style";
    import InputWithLabel
        from "../../../common/components/Inputs/InputWithLabel/InputWithLabel.svelte";
    import TextAreaWithLabel
        from "../../../common/components/Inputs/TextAreaWithLabel/TextAreaWithLabel.svelte";
    import InputDropdown
        from "../../../common/components/Inputs/InputDropdown/InputDropdown.svelte";

    const validateForm = (event: Event<HTMLInputElement>) => {
        console.log("validation")
    }

    const onShopSelect = (selectedId: number) => {
        console.log(selectedId)
    }

    onMount(() => {
        validateForm(null);
    })

    let purchase: Purchase = {
        shop: {},
        goods: [{}],
    };

</script>


<form class="{SideMinorPadding} {Form}">
    <div class="{FlexHor} {MainFieldsWrapper}">
        <div class="{FlexVert}" style="width: 50%; margin-right: 32px">
            <InputWithLabel
                    on:input={validateForm} on:change={validateForm}
                    label={$_("budget.labels.price")} autofocus={true}
                    type="number" name="totalPrice"
                    bind:value={purchase.totalPrice} required={true}/>
            <InputWithLabel
                    on:input={validateForm} on:change={validateForm}
                    label={$_("budget.labels.date")} autofocus={true}
                    type="date" name="date"
                    bind:value={purchase.date} required={true}/>
            <InputDropdown onSelectHandler={onShopSelect}
                           bind:value={purchase.shop.id} name="shop"
                           label={$_("budget.labels.shop")}
                           data={getSimpleShopsData()}/>
        </div>
        <div class="{FlexVert}" style="width: 50%">
            <TextAreaWithLabel
                    on:input={validateForm} on:change={validateForm}
                    textAreaClass={TextArea}
                    label={$_("common.labels.comment")} name="comment"
                    bind:value={purchase.comment}/>
        </div>
    </div>
    {#each purchase.goods as goodsItem (goodsItem.id)}
        <div class="{FlexHor}">
            <div class="{FlexVert}" style="width: 30%; margin-right: 32px">
                <InputWithLabel
                        on:input={validateForm} on:change={validateForm}
                        label={$_("budget.labels.price")} autofocus={true}
                        type="number" name="totalPrice"
                        bind:value={purchase.totalPrice} required={true}/>
            </div>
            <div class="{FlexVert}" style="width: 30%; margin-right: 32px">
            </div>
            <div class="{FlexVert}" style="width: 30%">
            </div>
        </div>
    {/each}
</form>