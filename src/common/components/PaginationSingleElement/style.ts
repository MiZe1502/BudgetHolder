
    import { css } from 'emotion';
    import colors from '../../styles/colors';
    
    export { FlexHorCenter } from "../../styles/positioning";
    export { Font316Black } from "../../styles/fonts";
    
    export const SingleButton = css`
        height: 25px;
        width: 35px;
        justify-content: center;
        border-radius: 4px;
        margin-right: 4px;
        padding-bottom: 2px;

        &:hover {
            cursor: pointer;
            background-color: ${colors.greenLight};
        }
    `

    export const ActiveButton = css`
       background-color: ${colors.greenActive};

       &:hover {
        background-color: ${colors.greenActive};
       }
    `