import DS from 'ember-data';

let {
    Model,
    attr
} = DS;

export default Model.extend({
  title: attr('string'),
  createdDate: attr('date'),
  description: attr('string')
});
